package postgres

import (
	"context"
	"main/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) *Repository {
	return &Repository{database: database}
}

func (s *Repository) GetPermissions(ctx context.Context, user *models.User) (*models.Permissions, error) {
	permissions := &models.Permissions{}
	err := s.database.QueryRowxContext(ctx, "SELECT JSONB_AGG(permission) permissions FROM users_role_permissions WHERE role = $1;", user.Role).Scan(permissions)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
