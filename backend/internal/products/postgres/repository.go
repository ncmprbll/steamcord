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
				SELECT
					products.id,
					products.name,
					products.discount,
					jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) AS price,
					products_images.tier_background_img
				FROM products
					JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) AS final FROM products_prices WHERE currency_code = $1) h
						ON products.id = h.product_id
					JOIN currencies ON currencies.code = h.currency_code
					JOIN products_images ON products.id = products_images.product_id
				GROUP BY products.id, products_images.tier_background_img, price, currencies.symbol, final
				ORDER BY RANDOM()
				LIMIT $2
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
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) AS price,
						products_images.tier_background_img
					FROM products
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) AS final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
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
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) AS price,
						products_images.featured_background_img,
						products_images.featured_logo_img,
						products_featured.created_at
					FROM products
						JOIN products_featured ON products.id = products_featured.product_id
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) AS final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
						JOIN currencies ON currencies.code = h.currency_code
						JOIN products_images ON products.id = products_images.product_id
					GROUP BY id, featured_background_img, featured_logo_img, price, currencies.symbol, final, products_featured.created_at
				), cart_items_platforms AS (
					SELECT
						id,
						name,
						discount,
						price,
						featured_background_img,
						featured_logo_img,
						created_at,
						jsonb_agg(products_platforms.platform) AS platforms
					FROM cart_items_price_image_featured
						JOIN products_platforms ON id = products_platforms.product_id
					GROUP BY id, name, discount, price, featured_background_img, featured_logo_img, created_at
				)
				SELECT
					id,
					name,
					discount,
					price,
					featured_background_img,
					featured_logo_img,
					platforms
				FROM cart_items_platforms
				ORDER BY created_at;
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
				SELECT COALESCE(json_agg(product_id), '[]'::json) AS products FROM users_games WHERE user_id = $1;
				`
	json := &models.JSONOwnedProducts{}
	if err := s.database.QueryRowxContext(ctx, query, user.UUID).Scan(json); err != nil {
		return nil, err
	}

	return json, nil
}

func (s *Repository) FindByID(ctx context.Context, product *models.Product, currencyCode, locale string) (*models.Product, error) {
	const query = `
				WITH product_price_screenshots AS (
					SELECT
						products.id,
						products.name,
						products.discount,
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) AS price,
						products_images.tier_background_img,
						COALESCE(jsonb_agg(products_screenshots.img) FILTER (WHERE products_screenshots.img IS NOT NULL), '[]'::jsonb) AS screenshots,
						about_token,
						description_token
					FROM products
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) AS final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
						JOIN currencies ON currencies.code = h.currency_code
						JOIN products_images ON products.id = products_images.product_id
						LEFT JOIN products_screenshots ON products.id = products_screenshots.product_id
					WHERE id = $2
					GROUP BY id, price, currencies.symbol, final, tier_background_img
				), product_platforms AS (
					SELECT
						id,
						name,
						discount,
						price,
						tier_background_img,
						screenshots,
						about_token,
						description_token,
						jsonb_agg(products_platforms.platform) AS platforms
					FROM product_price_screenshots
						JOIN products_platforms ON id = products_platforms.product_id
					GROUP BY id, name, discount, price, tier_background_img, screenshots, about_token, description_token
				), translated AS (
					SELECT
						id,
						name,
						discount,
						price,
						tier_background_img,
						screenshots,
						locale,
						MAX(CASE WHEN token = about_token THEN COALESCE(text, '') END) about,
						MAX(CASE WHEN token = description_token THEN COALESCE(text, '') END) description,
						platforms
					FROM translations
						RIGHT JOIN product_platforms ON (locale = $3 OR locale = 'en') AND (token = about_token OR token = description_token)
					GROUP BY id, name, discount, price, tier_background_img, screenshots, locale, platforms
				)
				SELECT
					*
				FROM translated;
				`
	rows, err := s.database.QueryxContext(ctx, query, currencyCode, product.ID, locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.Product{}

	for rows.Next() {
		row := &models.Product{}
		rows.StructScan(row)
		result = append(result, row)
	}

	for _, v := range result {
		if v.Locale.String == locale {
			return v, nil
		}
	}

	return result[0], nil
}

func (s *Repository) Search(ctx context.Context, currencyCode, name string, priceRange []float32, specials string, genres []string, pageLimit, pageOffset int) ([]*models.SearchProduct, error) {
	const baseQuery = `
				WITH products_price AS (
					SELECT
						products.id,
						products.name,
						products.discount,
						jsonb_build_object('original', h.price, 'final', h.final, 'symbol', currencies.symbol) AS price,
						products_images.tier_background_img,
						ARRAY_AGG(genres.genre) AS genres,
						created_at
					FROM products
						JOIN LATERAL (SELECT *, (price - (price * products.discount / 100)::NUMERIC(16, 2)) AS final FROM products_prices WHERE currency_code = $1) h ON products.id = h.product_id
						JOIN currencies ON currencies.code = h.currency_code
						JOIN products_images ON products.id = products_images.product_id
						LEFT JOIN products_genres ON products.id = products_genres.product_id
						LEFT JOIN genres ON products_genres.genre_id = genres.id
					WHERE LOWER(name) LIKE '%' || LOWER($2) || '%' AND h.final BETWEEN $3 AND $4 AND discount <> $5
					GROUP BY products.id, price, currencies.symbol, final, tier_background_img, created_at
				), product_platforms AS (
					SELECT
						id,
						name,
						discount,
						price,
						tier_background_img,
						jsonb_agg(products_platforms.platform) AS platforms,
						created_at
					FROM products_price
						JOIN products_platforms ON id = products_platforms.product_id
						WHERE LOWER(genres::TEXT)::TEXT[] @> LOWER($6::TEXT[]::TEXT)::TEXT[]
					GROUP BY id, name, discount, price, genres, tier_background_img, created_at
				)
				SELECT
					*
				FROM product_platforms
				ORDER BY created_at, id
				LIMIT $7 OFFSET $8;
				`
	result := []*models.SearchProduct{}

	rows, err := s.database.QueryxContext(ctx, baseQuery, currencyCode, name, priceRange[0], priceRange[1], specials, genres, pageLimit, pageOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := &models.SearchProduct{}
		rows.StructScan(row)
		result = append(result, row)
	}

	return result, nil
}
