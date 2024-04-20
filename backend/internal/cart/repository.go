package cart

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	Cart(context.Context, string, *models.User) ([]*models.CartGameRow, error)
	CartIDs(context.Context, *models.User) (*models.JSONCartProducts, error)
	Purchase(context.Context, *models.User) error
	AddToCart(context.Context, *models.CartRow) error
	DeleteFromCart(context.Context, *models.CartRow) (int64, error)
}
