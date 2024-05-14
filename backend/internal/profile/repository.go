package profile

import (
	"context"
	"main/backend/internal/models"
)

type Repository interface {
	Update(context.Context, *models.UserGeneralUpdate) error
	PasswordUpdate(context.Context, *models.UserPasswordUpdate) error
	PrivacyUpdate(context.Context, *models.UserPrivacyUpdate) error
	DeleteAvatar(context.Context, *models.User) (string, error)
	CreateComment(context.Context, *models.Comment) error
	GetComments(context.Context, *models.User, int, int) (*models.ProfileComments, error)
	DeleteComment(context.Context, *models.User, *models.User, *models.Comment) (bool, error)
	IsFriend(context.Context, *models.User, *models.User) (bool, error)
	HasIncomingInvite(context.Context, *models.User, *models.User) (bool, error)
	HasOutgoingInvite(context.Context, *models.User, *models.User) (bool, error)
	FriendInvite(context.Context, *models.User, *models.User) error
	HandleFriendInvite(context.Context, *models.User, *models.User, string) error
	DeleteFriend(context.Context, *models.User, *models.User) (bool, error)
}
