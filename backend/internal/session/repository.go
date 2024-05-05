package session

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	CreateSession(context.Context, *models.Session, int) (string, error)
	GetSessionByID(context.Context, string) (*models.Session, error)
	DeleteByID(context.Context, string) error
	InvalidateSessions(context.Context, *models.Session) error
}
