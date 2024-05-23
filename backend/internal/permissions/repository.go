package permissions

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetPermissions(context.Context, *models.User) (*models.Permissions, error)
}
