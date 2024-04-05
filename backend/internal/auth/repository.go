package auth

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	Register(context.Context, *models.User) error
}
