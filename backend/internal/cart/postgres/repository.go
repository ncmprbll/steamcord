package postgres

import (
	"context"
	"errors"
	"main/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) *Repository {
	return &Repository{database: database}
}

func (s *Repository) Cart(ctx context.Context, currencyCode string, user *models.User) ([]*models.CartGameRow, error) {
	const query = `
				WITH cart_ids AS (
					SELECT
						product_id
					FROM users_cart
					WHERE user_id = $2
				), cart_items_price_image AS (
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
					WHERE id IN (SELECT product_id FROM cart_ids)
					GROUP BY products.id, products_images.tier_background_img, price, currencies.symbol, final
				), cart_items_platforms AS (
					SELECT
						id,
						name,
						discount,
						price,
						tier_background_img,
						jsonb_agg(products_platforms.platform) as platforms
					FROM cart_items_price_image
						JOIN products_platforms ON id = products_platforms.product_id
					GROUP BY id, name, discount, price, tier_background_img
				)
				SELECT
					*
				FROM cart_items_platforms;
				`
	rows, err := s.database.QueryxContext(ctx, query, currencyCode, user.UUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*models.CartGameRow{}

	for rows.Next() {
		row := &models.CartGameRow{}
		rows.Scan(&row.ID, &row.Name, &row.Discount, &row.Price, &row.TierBackgroundImg, &row.Platforms)
		result = append(result, row)
	}

	return result, nil
}

func (s *Repository) CartIDs(ctx context.Context, user *models.User) (*models.JSONCartProducts, error) {
	const query = `
				SELECT COALESCE(json_agg(product_id), '[]'::json) as products FROM users_cart WHERE user_id = $1;
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

func (s *Repository) Purchase(ctx context.Context, user *models.User) error {
	tx, err := s.database.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	const queryTotal = `
						SELECT
							users.balance,
							SUM(price - (price * discount / 100)::NUMERIC(16, 2)) as total
						FROM users_cart
							JOIN users ON users_cart.user_id = $1
							JOIN products ON users_cart.product_id = products.id
							JOIN products_prices ON users_cart.product_id = products_prices.product_id
						WHERE products_prices.currency_code = users.currency_code GROUP BY users.balance;
						`
	var balance, total float32
	if err = tx.QueryRowxContext(ctx, queryTotal, user.UUID).Scan(&balance, &total); err != nil {
		return err
	}

	if balance < total {
		return errors.New("insufficient funds")
	}

	const queryBalance = `
						UPDATE
							users
						SET
							balance = balance - $1
						WHERE users.user_id = $2
						`

	if _, err = tx.ExecContext(ctx, queryBalance, total, user.UUID); err != nil {
		return err
	}

	const queryCart = `
						DELETE FROM users_cart
						WHERE users_cart.user_id = $1
						RETURNING product_id;
						`

	rows, err := tx.QueryxContext(ctx, queryCart, user.UUID)
	if err != nil {
		return err
	}
	defer rows.Close()

	const queryGames = `
						WITH operation AS (
							SELECT
								(price - (price * discount / 100)::NUMERIC(16, 2)) as spent
							FROM products
								JOIN products_prices ON products_prices.product_id = products.id
							WHERE id = $1 AND currency_code = $2
						)
						INSERT INTO users_games
							(user_id, product_id, currency_code, bought_for)
						VALUES
							($3, $1, $2, (SELECT spent FROM operation));
						`

	productIds := []int{}
	for rows.Next() {
		var productId int
		rows.Scan(&productId)
		productIds = append(productIds, productId)
	}

	for _, productId := range productIds {
		if _, err = tx.ExecContext(ctx, queryGames, productId, user.CurrencyCode, user.UUID); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
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
