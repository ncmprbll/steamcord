package cart

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetCartCount(context.Context, *models.User) (*models.JSONCartProducts, error)
	AddToCart(context.Context, *models.CartRow) error
}
