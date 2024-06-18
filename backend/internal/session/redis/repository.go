package redis

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"main/backend/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"crypto/sha256"
)

type Repository struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Repository {
	return &Repository{rdb: rdb}
}

func format(sessionId string) string {
	return fmt.Sprintf("session_id:%s", sessionId)
}

func SHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func (s *Repository) LogIPBadLoginAttempt(ctx context.Context, ip string, expiration int) error {
	attempts, err := s.rdb.Get(ctx, ip).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	err = s.rdb.Set(ctx, ip, attempts+1, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) GetIPBadLoginAttempts(ctx context.Context, ip string) (int, error) {
	attempts, err := s.rdb.Get(ctx, ip).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}

	return attempts, nil
}

func (s *Repository) CreateSession(ctx context.Context, session *models.Session, expiration int) (string, error) {
	sessionId := base64.StdEncoding.EncodeToString([]byte(uuid.New().String() + uuid.New().String()))

	sessionJson, err := json.Marshal(&session)
	if err != nil {
		return "", err
	}

	hash := SHA256(sessionId)

	err = s.rdb.Set(ctx, format(hash), sessionJson, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		return "", err
	}

	err = s.rdb.SAdd(ctx, format(session.UserID.String()), format(hash), 0).Err()
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (s *Repository) GetSessionByID(ctx context.Context, sessionId string) (*models.Session, error) {
	hash := SHA256(sessionId)

	bytes, err := s.rdb.Get(ctx, format(hash)).Bytes()
	if err != nil {
		return nil, err
	}

	session := &models.Session{}
	if err = json.Unmarshal(bytes, &session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Repository) DeleteSession(ctx context.Context, session *models.Session) error {
	hash := SHA256(session.SessionID)

	if err := s.rdb.Del(ctx, format(hash)).Err(); err != nil {
		return err
	}

	if err := s.rdb.SRem(ctx, format(session.UserID.String()), format(hash)).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Repository) InvalidateSessions(ctx context.Context, session *models.Session) error {
	setKey := format(session.UserID.String())
	setMembers := s.rdb.SMembers(context.TODO(), setKey)
	err := setMembers.Err()
	if err != nil {
		return err
	}

	sessions, err := setMembers.Result()
	if err != nil {
		return err
	}

	for _, v := range sessions {
		if err := s.rdb.Del(ctx, v).Err(); err != nil {
			return err
		}
	}

	if err := s.rdb.Del(ctx, setKey).Err(); err != nil {
		return err
	}

	return nil
}
