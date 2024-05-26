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
	MIN_PASSWORD_LENGTH = 8
	MAX_PASSWORD_LENGTH = 48
	MIN_LOGIN_LENGTH    = 6
	MAX_LOGIN_LENGTH    = 20

	MIN_DISPLAY_NAME_LENGTH = 1
	MAX_DISPLAY_NAME_LENGTH = 20
	MAX_ABOUT_LENGTH        = 256
)

type User struct {
	UUID         uuid.UUID  `json:"id" db:"id"`
	Login        string     `json:"login,omitempty" db:"login"`
	Avatar       string     `json:"avatar" db:"avatar"`
	DisplayName  string     `json:"display_name" db:"display_name"`
	About        string     `json:"about,omitempty" db:"about"`
	Privacy      string     `json:"privacy,omitempty" db:"privacy"`
	CurrencyCode string     `json:"currency_code,omitempty" db:"currency_code"`
	Balance      *float32   `json:"balance,omitempty" db:"balance"`
	Email        string     `json:"email,omitempty" db:"email"`
	Password     string     `json:"password,omitempty" db:"password"`
	Role         string     `json:"role,omitempty" db:"role"`
	Banned       *bool      `json:"banned,omitempty" db:"banned"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	LoginDate    *time.Time `json:"login_date,omitempty" db:"login_date"`

	Hidden bool `json:"hidden,omitempty"`
}

type PublicUser struct {
	UUID        uuid.UUID `json:"id" db:"id"`
	Avatar      string    `json:"avatar" db:"avatar"`
	DisplayName string    `json:"display_name" db:"display_name"`
}

type UserGeneralUpdate struct {
	UUID        uuid.UUID `json:"id" db:"id"`
	Avatar      string    `json:"avatar" db:"avatar"`
	DisplayName string    `json:"display_name" db:"display_name"`
	About       string    `json:"about" db:"about"`
}

type UserPasswordUpdate struct {
	UUID        uuid.UUID `json:"id"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

type UserPrivacyUpdate struct {
	UUID    uuid.UUID `json:"id"`
	Privacy string    `json:"privacy"`
}

type ManagementUsers struct {
	Users      []*User  `json:"users"`
	Total      int      `json:"total"`
	Roles      []string `json:"roles"`
	Currencies []string `json:"currencies"`
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

func (u *User) RemoveSensitiveData() {
	u.Login = ""
	u.Email = ""
	u.Privacy = ""
	u.CurrencyCode = ""
	u.Balance = nil
	u.SanitizePassword()
	u.Role = ""
	u.Banned = nil
	u.UpdatedAt = nil
}

func (u *User) Validate() error {
	login := u.Login
	email := u.Email
	password := u.Password

	if len(login) < MIN_LOGIN_LENGTH {
		return errors.New("validation error: login too short")
	} else if len(login) > MAX_LOGIN_LENGTH {
		return errors.New("validation error: login too long")
	}

	if len(password) < MIN_PASSWORD_LENGTH {
		return errors.New("validation error: password too short")
	} else if len(password) > MAX_PASSWORD_LENGTH {
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

func (u *User) ApplyPrivacy() {
	u.About = ""
	u.CreatedAt = nil
	u.LoginDate = nil
	u.Hidden = true
}

func (u *UserGeneralUpdate) Sanitize() {
	str := strings.TrimSpace(u.DisplayName)
	pattern := regexp.MustCompile(`\s+`)

	u.DisplayName = pattern.ReplaceAllString(str, " ")
	u.About = strings.TrimSpace(strings.ReplaceAll(u.About, "\r", ""))
}

func (u *UserGeneralUpdate) Validate() error {
	displayName := u.DisplayName
	about := u.About

	if displayName != "" {
		if len(displayName) < MIN_DISPLAY_NAME_LENGTH {
			return errors.New("validation error: displayName too short")
		} else if len(displayName) > MAX_DISPLAY_NAME_LENGTH {
			return errors.New("validation error: displayName too long")
		}
	}

	if len(about) > MAX_ABOUT_LENGTH {
		return errors.New("validation error: about too long")
	}

	return nil
}

func (u *UserPasswordUpdate) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.NewPassword = string(hashedPassword)
	return nil
}

func (u *UserPasswordUpdate) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.OldPassword)); err != nil {
		return err
	}
	return nil
}

func (u *UserPasswordUpdate) Validate() error {
	password := u.NewPassword

	if len(password) < MIN_PASSWORD_LENGTH {
		return errors.New("validation error: password too short")
	} else if len(password) > MAX_PASSWORD_LENGTH {
		return errors.New("validation error: password too long")
	}

	return nil
}
