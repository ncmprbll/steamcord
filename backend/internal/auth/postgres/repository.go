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

func (s *Repository) Register(ctx context.Context, user *models.User) error {
	_, err := s.database.ExecContext(ctx, queryRegisterUser, user.Login, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) FindByLogin(ctx context.Context, user *models.User) (*models.User, error) {
	found := &models.User{}
	if err := s.database.QueryRowxContext(ctx, "SELECT * FROM users WHERE login = $1;", user.Login).StructScan(found); err != nil {
		return nil, err
	}
	return found, nil
}

func (s *Repository) FindByUUID(ctx context.Context, user *models.User) (*models.User, error) {
	found := &models.User{}
	if err := s.database.QueryRowxContext(ctx, "SELECT * FROM users WHERE id = $1;", user.UUID).StructScan(found); err != nil {
		return nil, err
	}
	return found, nil
}

func (s *Repository) UpdateLoginDate(ctx context.Context, user *models.User) error {
	_, err := s.database.ExecContext(ctx, "UPDATE users SET login_date = NOW() WHERE id = $1;", user.UUID)
	if err != nil {
		return err
	}
	return nil
}
