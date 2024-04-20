package models

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type JSONProductID struct {
	ProductID int `json:"product_id"`
}

type JSONCartProducts []int

type CartRow struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	ProductID int     `json:"product_id,omitempty"`
}

type CartGameRow struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Discount          int           `json:"discount"`
	Price            JSONPrice     `json:"price"`
	TierBackgroundImg string        `json:"tier_background_img"`
	Platforms         JSONPlatforms `json:"platforms"`
}

func (p *JSONCartProducts) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}
