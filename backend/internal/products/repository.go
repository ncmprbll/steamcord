package products

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetTier(context.Context, string) ([]*models.GetTierRow, error)
	GetTierByGenre(context.Context, string, string) ([]*models.GetTierRow, error)
	GetFeatured(context.Context) ([]*models.GetFeaturedRow, error)
}
