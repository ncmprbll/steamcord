package models

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type JSONProductID struct {
	ProductID int64 `json:"product_id"`
}

type JSONCartProducts []int

type CartRow struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	ProductID int64     `json:"product_id,omitempty"`
}

func (p *JSONCartProducts) Scan(src any) error {
	bytes, ok := src.([]byte)

	if !ok {
		return errors.New("not a bytes array")
	}

	return json.Unmarshal(bytes, p)
}
