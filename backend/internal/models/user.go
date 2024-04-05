package models

import (
	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID `json:"user_id" db:"user_id"`
	Login string `json:"login" db:"login"`
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role string `json:"role" db:"role"`
}