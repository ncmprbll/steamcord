package session

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	LogIPBadLoginAttempt(context.Context, string, int) error
	GetIPBadLoginAttempts(context.Context, string) (int, error)
	CreateSession(context.Context, *models.Session, int) (string, error)
	GetSessionByID(context.Context, string) (*models.Session, error)
	DeleteSession(context.Context, *models.Session) error
	InvalidateSessions(context.Context, *models.Session) error
}
