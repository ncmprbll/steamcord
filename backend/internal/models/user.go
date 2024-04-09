package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	MIN_PASSWORD = 8
	MAX_PASSWORD = 48
	MIN_LOGIN = 6
	MAX_LOGIN = 20
)

type User struct {
	UUID        uuid.UUID `json:"user_id" db:"user_id"`
	Login       string    `json:"login" db:"login"`
	DisplayName string    `json:"display_name" db:"display_name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	LoginDate   time.Time `json:"login_date" db:"login_date"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

func (u *User) Validate() error {
	login := u.Login
	email := u.Email
	password := u.Password

	if len(login) < MIN_LOGIN {
		return errors.New("validation error: login too short")
	} else if len(login) > MAX_LOGIN {
		return errors.New("validation error: login too long")
	}

	if len(password) < MIN_PASSWORD {
		return errors.New("validation error: password too short")
	} else if len(password) > MAX_PASSWORD {
		return errors.New("validation error: password too long")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(login) {
		return errors.New("validation error: illegal login characters")
	}

	email = strings.ReplaceAll(email, " ", "")

	if email != u.Email {
		return errors.New("validation error: email space characters")
	}

	return nil
}
