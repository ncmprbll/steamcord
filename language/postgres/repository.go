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

func (s *Repository) GetAll(ctx context.Context) ([]*models.Locale, error) {
	rows, err := s.database.QueryxContext(ctx, "SELECT * FROM locales;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.Locale{}

	for rows.Next() {
		row := &models.Locale{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return result, nil
}
