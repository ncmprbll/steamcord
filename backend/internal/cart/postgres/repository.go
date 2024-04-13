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

func (s *Repository) AddToCart(ctx context.Context, cart *models.CartRow) error {
	const query = `
				INSERT INTO users_cart (user_id, product_id) VALUES ($1, $2);
				`
	_, err := s.database.ExecContext(ctx, query, cart.UserID, cart.ProductID)
	if err != nil {
		return err
	}

	return nil
}
