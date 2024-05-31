package products

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetTier(context.Context, string, string) ([]*models.TierRow, error)
	GetTierByGenre(context.Context, string, string, string) ([]*models.TierRow, error)
	GetFeatured(context.Context, string) ([]*models.FeaturedRow, error)
	GetOwnedIDs(context.Context, *models.User) (*models.JSONOwnedProducts, error)
	FindByID(context.Context, *models.Product, string, string) (*models.Product, error)
	Search(context.Context, string, string, []float32, string, []string, int, int) ([]*models.SearchProduct, error)
	Currencies(context.Context) (*models.Currencies, error)
	CreateProduct(context.Context, *models.PublishProduct) error
}
