package models

import (
	"encoding/json"
	"errors"
)

type Price struct {
	ID           int     `json:"id,omitempty"`
	ProductID    int64   `json:"product_id,omitempty"`
	CurrencyCode string  `json:"currency_code,omitempty"`
	Price        float32 `json:"price,omitempty"`
}

type JSONPrice struct {
	Origianl float32 `json:"original"`
	Final    float32 `json:"final"`
	Symbol   string  `json:"symbol"`
}

type JSONPlatforms []string

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
