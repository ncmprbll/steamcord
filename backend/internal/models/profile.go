package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	MAX_COMMENT_LENGTH = 128
)

type Comment struct {
	ID          int       `json:"id,omitempty" db:"id"`
	ProfileID   string    `json:"profile_id,omitempty" db:"profile_id"`
	Commnetator uuid.UUID `json:"commentator" db:"commentator"`
	Text        string    `json:"text" db:"text"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
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
