package products

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetTier(context.Context, int) ([]*models.GetTierRow, error)
	GetFeatured(context.Context) ([]*models.GetFeaturedRow, error)
}
