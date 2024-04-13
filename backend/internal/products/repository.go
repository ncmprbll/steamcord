package products

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetTier(context.Context, string) ([]*models.TierRow, error)
	GetTierByGenre(context.Context, string, string) ([]*models.TierRow, error)
	GetFeatured(context.Context) ([]*models.FeaturedRow, error)
}
