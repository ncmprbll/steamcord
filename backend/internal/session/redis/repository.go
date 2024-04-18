package redis

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"main/backend/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
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

func (s *Repository) CreateSession(ctx context.Context, session *models.Session, expiration int) (string, error) {
	sessionId := base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))

	sessionJson, err := json.Marshal(&session)
	if err != nil {
		return "", err
	}

	err = s.rdb.Set(ctx, format(sessionId), sessionJson, time.Duration(expiration) * time.Second).Err()
	if err != nil {
		return "", err
	}

	return sessionId, err
}

func (s *Repository) GetSessionByID(ctx context.Context, sessionId string) (*models.Session, error) {
	bytes, err := s.rdb.Get(ctx, format(sessionId)).Bytes()
	if err != nil {
		return nil, err
	}

	session := &models.Session{}
	if err = json.Unmarshal(bytes, &session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Repository) DeleteByID(ctx context.Context, sessionId string) error {
	if err := s.rdb.Del(ctx, format(sessionId)).Err(); err != nil {
		return err
	}

	return nil
}
