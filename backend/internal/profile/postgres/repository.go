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

func (s *Repository) Update(ctx context.Context, user *models.UserGeneralUpdate) error {
	const query = `
				UPDATE
					users
				SET
					avatar = CASE
						WHEN $1 = '' THEN avatar
						WHEN $1 <> '' THEN $1
					END,
					display_name = CASE
						WHEN $2 = '' THEN display_name
						WHEN $2 <> '' THEN $2
					END,
					about = $3
				WHERE user_id = $4;
				`
	_, err := s.database.ExecContext(ctx, query, user.Avatar, user.DisplayName, user.About, user.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) PasswordUpdate(ctx context.Context, user *models.UserPasswordUpdate) error {
	const query = `
				UPDATE
					users
				SET
					password = $1
				WHERE user_id = $2;
				`
	_, err := s.database.ExecContext(ctx, query, user.NewPassword, user.UUID)
	if err != nil {
		return err
	}

	return nil
}
