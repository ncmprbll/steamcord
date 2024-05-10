package postgres

import (
	"context"
	"main/backend/internal/models"

	"github.com/google/uuid"
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
				WHERE id = $4;
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
				WHERE id = $2;
				`
	_, err := s.database.ExecContext(ctx, query, user.NewPassword, user.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) PrivacyUpdate(ctx context.Context, user *models.UserPrivacyUpdate) error {
	const query = `
				UPDATE
					users
				SET
					privacy = $1
				WHERE id = $2;
				`
	_, err := s.database.ExecContext(ctx, query, user.Privacy, user.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Repository) DeleteAvatar(ctx context.Context, user *models.User) (string, error) {
	const query = `
				WITH u AS (
					SELECT avatar FROM users WHERE id = $1
				)
				UPDATE
					users
				SET
					avatar = ''
				WHERE id = $1
				RETURNING (SELECT avatar FROM u);
				`
	var avatar string
	if err := s.database.QueryRowxContext(ctx, query, user.UUID).Scan(&avatar); err != nil {
		return "", err
	}

	return avatar, nil
}

func (s *Repository) CreateComment(ctx context.Context, comment *models.Comment) error {
	const query = `
				INSERT INTO
					users_comments (profile_id, commentator, text)
				VALUES ($1, $2, $3);
				`
	_, err := s.database.ExecContext(ctx, query, comment.ProfileID, comment.Commnetator, comment.Text)
	if err != nil {
		return err
	}
	return nil
}

func (s *Repository) GetComments(ctx context.Context, uuid uuid.UUID, pageLimit, pageOffset int) ([]*models.Comment, error) {
	const query = `
				SELECT
					commentator,
					text,
					created_at
				FROM users_comments
				WHERE profile_id = $1
				ORDER BY created_at DESC
				LIMIT $2 OFFSET $3;
				`
	rows, err := s.database.QueryxContext(ctx, query, uuid, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.Comment{}

	for rows.Next() {
		row := &models.Comment{}
		rows.StructScan(&row)
		result = append(result, row)
	}

	return result, nil
}
