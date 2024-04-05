package postgres

import (
	"context"
	"database/sql"
	"main/backend/internal/models"
)

type Repository struct {
	database *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{database: database}
}

func (s *Repository) Register(ctx context.Context, user *models.User) error {
	_, err := s.database.ExecContext(ctx, "INSERT INTO users (login, email, password) VALUES ($1, $2, $3);", user.Login, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}