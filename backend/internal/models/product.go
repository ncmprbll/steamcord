package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

const (
	PERMISSION_UI_PUBLISHING = "ui.publishing"

	PRODUCTS_PAGE_LIMIT = 15
)

type JSONOwnedProducts []int

type Price struct {
	ProductID    int64   `json:"product_id" db:"product_id"`
	CurrencyCode string  `json:"currency_code" db:"currency_code"`
	Price        float32 `json:"price" db:"price"`
}

type JSONPrice struct {
	Origianl float32 `json:"original"`
	Final    float32 `json:"final"`
	Symbol   string  `json:"symbol"`
}

type JSONPlatforms []string
type Screenshots []string

type TierRow struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Discount          int       `json:"discount"`
	Price             JSONPrice `json:"price"`
	TierBackgroundImg string    `json:"tier_background_img"`
}

type FeaturedRow struct {
	ID                    int           `json:"id"`
	Name                  string        `json:"name"`
	Discount              int           `json:"discount"`
	Price                 JSONPrice     `json:"price"`
	FeaturedBackgroundImg string        `json:"featured_background_img"`
	FeaturedLogoImg       string        `json:"featured_logo_img"`
	Platforms             JSONPlatforms `json:"platforms"`
}

type Product struct {
	ID                int            `json:"id" db:"id"`
	Name              string         `json:"name" db:"name"`
	Discount          int            `json:"discount" db:"discount"`
	Publisher         string         `json:"publisher" db:"publisher"`
	Price             JSONPrice      `json:"price" db:"price"`
	TierBackgroundImg string         `json:"tier_background_img" db:"tier_background_img"`
	Screenshots       Screenshots    `json:"screenshots" db:"screenshots"`
	Locale            sql.NullString `json:"-" db:"locale"`
	About             sql.NullString `json:"about" db:"about"`
	Description       sql.NullString `json:"description" db:"description"`
	Platforms         JSONPlatforms  `json:"platforms" db:"platforms"`
	CreatedAt         *time.Time     `json:"created_at,omitempty" db:"created_at"`
}

type SearchProduct struct {
	ID                int           `json:"id" db:"id"`
	Name              string        `json:"name" db:"name"`
	Discount          int           `json:"discount" db:"discount"`
	Price             JSONPrice     `json:"price" db:"price"`
	TierBackgroundImg string        `json:"tier_background_img" db:"tier_background_img"`
	Platforms         JSONPlatforms `json:"platforms" db:"platforms"`
	CreatedAt         time.Time     `json:"-" db:"created_at"`
}

type Currency struct {
	Code   string `json:"code" db:"code"`
	Symbol string `json:"symbol" db:"symbol"`
}

type Genre struct {
	Id    string `json:"id" db:"id"`
	Genre string `json:"genre" db:"genre"`
}

type Currencies []*Currency

type PublishProduct struct {
	Name        string             `json:"name"`
	Header      string             `json:"header"`
	Screenshots []string           `json:"screenshots"`
	About       map[string]string  `json:"about"`
	Description map[string]string  `json:"description"`
	Prices      map[string]float32 `json:"prices"`
}

type UpdateProduct struct {
	ID       int                `json:"id"`
	Discount int                `json:"discount"`
	Prices   map[string]float32 `json:"prices"`
}

type Sales []*struct {
	Date  string `json:"date" db:"date"`
	Sales string `json:"sales" db:"sales"`
}

func (p *JSONPrice) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}

func (p *JSONPlatforms) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}

func (p *JSONOwnedProducts) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}

func (p *Screenshots) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}
