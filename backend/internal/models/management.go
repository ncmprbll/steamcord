package models

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	PERMISSION_UI_MANAGEMENT    = "ui.management"
	PERMISSION_USERS_MANAGEMENT = "management.users"
	PERMISSION_ROLES_MANAGEMENT = "management.roles"
)

type Permission struct {
	Role       string     `json:"role" db:"role"`
	Permission string     `json:"permission" db:"permission"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Permissions []string

type Role struct {
	Name      string     `json:"name" db:"name"`
	CanDelete bool       `json:"can_delete" db:"can_delete"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type Roles []*Role

func (p *Permissions) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}
