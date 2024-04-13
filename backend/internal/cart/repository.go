package cart

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	AddToCart(context.Context, *models.CartRow) error
}
