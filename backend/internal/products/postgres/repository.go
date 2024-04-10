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
	const query = `
				SELECT products.id, products.name, products.discount, json_object_agg(products_prices.currency_code, products_prices.price) as prices, products_images.tier_background_img
					FROM products
						JOIN products_prices ON products.id = products_prices.product_id
						JOIN products_images ON products.id = products_images.product_id
					WHERE products_images.tier_background_img <> ''
					GROUP BY products.id, products_images.tier_background_img
					ORDER BY RANDOM()
					LIMIT $1;
				`
	rows, err := s.database.QueryxContext(ctx, query, limit)
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

func (s *Repository) GetFeatured(ctx context.Context) ([]*models.GetFeaturedRow, error) {
	const query = `
				SELECT products.id, products.name, products.discount, json_object_agg(products_prices.currency_code, products_prices.price) as prices, products_images.featured_background_img, products_images.featured_logo_img, json_agg(products_platforms.platform) as platforms
					FROM products
						JOIN products_featured ON products.id = products_featured.product_id
						JOIN products_prices ON products.id = products_prices.product_id
						JOIN products_images ON products.id = products_images.product_id
						JOIN products_platforms ON products.id = products_platforms.product_id
					GROUP BY products.id, products_images.featured_background_img, products_images.featured_logo_img;
				`
	rows, err := s.database.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.GetFeaturedRow{}

	for rows.Next() {
		row := &models.GetFeaturedRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Prices, &row.FeaturedBackgroundImg, &row.FeaturedLogoImg, &row.Platforms)
		result = append(result, row)
	}

	return result, nil
}
