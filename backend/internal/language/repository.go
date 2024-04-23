package language

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	GetAll(context.Context) ([]*models.Locale, error)
}
