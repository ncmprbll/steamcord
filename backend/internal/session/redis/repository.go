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

func (s *Repository) CreateSession(ctx context.Context, session *models.Session, expiration int) (string, error) {
	token := base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))

	sessionJson, err := json.Marshal(&session)
	if err != nil {
		return "", err
	}

	err = s.rdb.Set(ctx, fmt.Sprintf("session_id::%s", token), sessionJson, time.Duration(expiration) * time.Second).Err()
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *Repository) GetSessionByID(ctx context.Context, sessionId string) (*models.Session, error) {
	return nil, nil
}

func (s *Repository) DeleteByID(ctx context.Context, sessionId string) error {
	return nil
}
