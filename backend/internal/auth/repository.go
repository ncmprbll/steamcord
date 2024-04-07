package auth

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	Register(context.Context, *models.User) error
	FindByLogin(context.Context, *models.User) (*models.User, error)
	FindByUUID(context.Context, *models.User) (*models.User, error)
}
