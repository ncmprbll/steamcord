package cart

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	Cart(context.Context, *models.User) ([]*models.CartGameRow, error)
	CartCount(context.Context, *models.User) (*models.JSONCartProducts, error)
	AddToCart(context.Context, *models.CartRow) error
}
