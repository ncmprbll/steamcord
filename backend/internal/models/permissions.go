package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Permission struct {
	Role       string     `json:"role" db:"role"`
	Permission string     `json:"permission" db:"permission"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Permissions []string

func (p *Permissions) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}
