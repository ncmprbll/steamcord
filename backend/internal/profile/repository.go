package profile

import (
	"context"
	"main/backend/internal/models"

	"github.com/google/uuid"
)

type Repository interface {
	Update(context.Context, *models.UserGeneralUpdate) error
	PasswordUpdate(context.Context, *models.UserPasswordUpdate) error
	DeleteAvatar(context.Context, *models.User) (string, error)
	CreateComment(context.Context, *models.Comment) error
	GetComments(context.Context, uuid.UUID, int, int) ([]*models.Comment, error)
}
