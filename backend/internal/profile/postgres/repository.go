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

func (s *Repository) Update(ctx context.Context, user *models.User) error {
	const query = `
				UPDATE
					users
				SET
					display_name = CASE
						WHEN $1 = '' THEN display_name
						WHEN $1 <> '' THEN $1
					END,
					about = CASE
						WHEN $2 = '' THEN about
						WHEN $2 <> '' THEN $2
					END
				WHERE user_id = $3;
				`
	_, err := s.database.ExecContext(ctx, query, user.DisplayName, user.About, user.UUID)
	if err != nil {
		return err
	}

	return nil
}
