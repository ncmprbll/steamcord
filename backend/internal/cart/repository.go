package cart

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetCart(context.Context) 
	AddToCart(context.Context, *models.CartRow) error
}
