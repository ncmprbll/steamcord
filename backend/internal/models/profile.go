package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	FRIEND_REQUEST_REJECTED_WAIT_TIME = time.Hour * 24
	COMMENTS_PAGE_LIMIT               = 10
)

var (
	ErrCannotFriendSelf             = errors.New("cannot befriend yourself")
	ErrAlreadyFriends               = errors.New("already friends")
	ErrFriendInvitePending          = errors.New("an invite is already pending")
	ErrAlreadyTriedToFriendRejected = errors.New("friend request has been rejected, try again later")
	ErrNoFriendRequest              = errors.New("no friend request")
)

const (
	MAX_COMMENT_LENGTH = 128
)

type Comment struct {
	ID          int       `json:"id,omitempty" db:"id"`
	ProfileID   string    `json:"profile_id,omitempty" db:"profile_id"`
	Commentator uuid.UUID `json:"commentator" db:"commentator"`
	Text        string    `json:"text" db:"text"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type DisplayComment struct {
	ID          int       `json:"id,omitempty" db:"id"`
	Commentator uuid.UUID `json:"commentator" db:"commentator"`
	Avatar      string    `json:"avatar" db:"avatar"`
	DisplayName string    `json:"display_name" db:"display_name"`
	Text        string    `json:"text" db:"text"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type ProfileComments struct {
	Comments []*DisplayComment `json:"comments"`
	Total    int               `json:"total"`
}

type Invite struct {
	ID        int       `json:"id,omitempty" db:"id"`
	Invitee   uuid.UUID `json:"invitee" db:"invitee"`
	Inviter   uuid.UUID `json:"inviter" db:"inviter"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type FriendStatus struct {
	IsFriend          bool `json:"isFriend"`
	HasIncomingInvite bool `json:"hasIncomingInvite"`
	HasOutgoingInvite bool `json:"hasOutgoingInvite"`
}

type Games struct {
	Games []*Product `json:"games"`
	Total int        `json:"total"`
}

func (c *Comment) Validate() error {
	text := c.Text

	if len(text) > MAX_COMMENT_LENGTH {
		return errors.New("validation error: login too short")
	}

	return nil
}

func (c *Comment) Sanitize() {
	str := strings.TrimSpace(c.Text)
	pattern := regexp.MustCompile(`\s+`)

	c.Text = pattern.ReplaceAllString(str, " ")
	c.Text = strings.ReplaceAll(c.Text, "\r", "")
	c.Text = strings.ReplaceAll(c.Text, "\n", "")
}
