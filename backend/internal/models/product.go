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

type JSONPrice map[string]float32
type JSONPlatforms []string

type GetTierRow struct {
	ID                int       `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Discount          int       `json:"discount,omitempty"`
	Prices            JSONPrice `json:"prices,omitempty"`
	TierBackgroundImg string    `json:"tier_background_img,omitempty"`
}

type GetFeaturedRow struct {
	ID                    int           `json:"id,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Discount              int           `json:"discount,omitempty"`
	Prices                JSONPrice     `json:"prices,omitempty"`
	FeaturedBackgroundImg string        `json:"featured_background_img,omitempty"`
	FeaturedLogoImg       string        `json:"featured_logo_img,omitempty"`
	Platforms             JSONPlatforms `json:"platforms,omitempty"`
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
