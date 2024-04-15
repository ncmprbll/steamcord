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

func (s *Repository) Cart(ctx context.Context, user *models.User) ([]*models.CartGameRow, error) {
	const query = `
				SELECT products.id, products.name, products.discount, json_object_agg(products_prices.currency_code, products_prices.price) as prices, products_images.tier_background_img, json_agg(products_platforms.platform) as platforms
				FROM products
					JOIN products_prices ON products.id = products_prices.product_id
					JOIN products_images ON products.id = products_images.product_id
					JOIN products_platforms ON products.id = products_platforms.product_id
				WHERE id IN (SELECT product_id FROM users_cart WHERE user_id = $1)
				GROUP BY products.id, products_images.tier_background_img;
				`
	rows, err := s.database.QueryxContext(ctx, query, user.UUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.CartGameRow{}

	for rows.Next() {
		row := &models.CartGameRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Prices, &row.TierBackgroundImg, &row.Platforms)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) CartIDs(ctx context.Context, user *models.User) (*models.JSONCartProducts, error) {
	const query = `
				SELECT json_agg(product_id) as products FROM users_cart WHERE user_id = $1;
				`
	json := &models.JSONCartProducts{}
	if err := s.database.QueryRowxContext(ctx, query, user.UUID).Scan(json); err != nil {
		return nil, err
	}

	return json, nil
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

func (s *Repository) DeleteFromCart(ctx context.Context, cart *models.CartRow) (int64, error) {
	const query = `
				DELETE FROM users_cart WHERE user_id = $1 AND product_id = $2;
				`
	result, err := s.database.ExecContext(ctx, query, cart.UserID, cart.ProductID)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}
