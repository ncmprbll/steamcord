package profile

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	Update(context.Context, *models.User) error
}
