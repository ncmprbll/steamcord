package management

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetPermissions(context.Context, *models.User) (*models.Permissions, error)
	GetUsers(context.Context, string) (*models.ManagementUsers, error)
	UpdateUser(context.Context, *models.User) error
	GetRoles(context.Context) (*models.Roles, error)
	CreateRole(context.Context, *models.Role) error
	DeleteRole(context.Context, *models.Role) (int64, error)
	GetRolePermissions(context.Context) (*models.RolePermissions, error)
	AddPermission(context.Context, *models.Role, *models.Permissions) error
	DeletePermission(context.Context, *models.Role, *models.Permissions) (int64, error)
}
