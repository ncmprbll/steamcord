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

func (s *Repository) GetTier(ctx context.Context, currencyCode, limit string) ([]*models.TierRow, error) {
	const query = `
				WITH cart_items_price_image AS (
					SELECT
						products.id,
						products.name,
						products.discount,
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) as price,
						products_images.tier_background_img
					FROM products
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) as final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
						JOIN currencies ON currencies.code = h.currency_code
						JOIN products_images ON products.id = products_images.product_id
					GROUP BY products.id, products_images.tier_background_img, price, currencies.symbol, final
					ORDER BY RANDOM()
					LIMIT $2
				)
				SELECT
					*
				FROM cart_items_price_image;
				`
	rows, err := s.database.QueryxContext(ctx, query, currencyCode, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.TierRow{}

	for rows.Next() {
		row := &models.TierRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Price, &row.TierBackgroundImg)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetTierByGenre(ctx context.Context, currencyCode, genre, limit string) ([]*models.TierRow, error) {
	const query = `
				WITH cart_items_price_image AS (
					SELECT
						products.id,
						products.name,
						products.discount,
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) as price,
						products_images.tier_background_img
					FROM products
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) as final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
						JOIN currencies ON currencies.code = h.currency_code
						JOIN products_images ON products.id = products_images.product_id
						JOIN products_genres ON products.id = products_genres.product_id
						JOIN genres ON products_genres.genre_id = genres.id
					GROUP BY products.id, products_images.tier_background_img, price, currencies.symbol, final
					HAVING $2 = ANY(ARRAY_AGG(genres.genre))
					ORDER BY RANDOM()
					LIMIT $3
				)
				SELECT
					*
				FROM cart_items_price_image;
				`
	rows, err := s.database.QueryxContext(ctx, query, currencyCode, genre, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.TierRow{}

	for rows.Next() {
		row := &models.TierRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Price, &row.TierBackgroundImg)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetFeatured(ctx context.Context, currencyCode string) ([]*models.FeaturedRow, error) {
	const query = `
				WITH cart_items_price_image_featured AS (
					SELECT
						products.id,
						products.name,
						products.discount,
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) as price,
						products_images.featured_background_img,
						products_images.featured_logo_img
					FROM products
						JOIN products_featured ON products.id = products_featured.product_id
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) as final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
						JOIN currencies ON currencies.code = h.currency_code
						JOIN products_images ON products.id = products_images.product_id
					GROUP BY id, featured_background_img, featured_logo_img, price, currencies.symbol, final
				), cart_items_platforms AS (
					SELECT
						id,
						name,
						discount,
						price,
						featured_background_img,
						featured_logo_img,
						jsonb_agg(products_platforms.platform) as platforms
					FROM cart_items_price_image_featured
						JOIN products_platforms ON id = products_platforms.product_id
					GROUP BY id, name, discount, price, featured_background_img, featured_logo_img
				)
				SELECT
					*
				FROM cart_items_platforms;
				`
	rows, err := s.database.QueryxContext(ctx, query, currencyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.FeaturedRow{}

	for rows.Next() {
		row := &models.FeaturedRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Price, &row.FeaturedBackgroundImg, &row.FeaturedLogoImg, &row.Platforms)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) GetOwnedIDs(ctx context.Context, user *models.User) (*models.JSONOwnedProducts, error) {
	const query = `
				SELECT COALESCE(json_agg(product_id), '[]'::json) as products FROM users_games WHERE user_id = $1;
				`
	json := &models.JSONOwnedProducts{}
	if err := s.database.QueryRowxContext(ctx, query, user.UUID).Scan(json); err != nil {
		return nil, err
	}

	return json, nil
}