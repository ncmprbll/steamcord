package management

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetPermissions(context.Context, *models.User) (*models.Permissions, error)
	GetUsers(context.Context, string) (*models.ManagementUsers, error)
	UpdateUser(context.Context, *models.User) error
}
