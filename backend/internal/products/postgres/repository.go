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

func (s *Repository) GetTier(ctx context.Context, limit int) ([]*models.GetTierRow, error) {
	rows, err := s.database.QueryxContext(ctx, "SELECT products.id, products.name, products.discount, json_object_agg(products_prices.currency_code, products_prices.price) as prices, products_images.tier_background_img FROM products JOIN products_prices ON products.id = products_prices.product_id JOIN products_images ON products.id = products_images.product_id WHERE products_images.tier_background_img <> '' GROUP BY products.id, products_images.tier_background_img LIMIT $1;", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.GetTierRow{}

	for rows.Next() {
		row := &models.GetTierRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Prices, &row.TierBackgroundImg)
		result = append(result, row)
	}

	return result, nil
}