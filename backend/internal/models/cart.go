package models

import "github.com/google/uuid"

type JSONProductID struct {
	ProductID int64 `json:"product_id"`
}

type CartRow struct {
	UserID    uuid.UUID `json:"user_id"`
	ProductID int64     `json:"product_id"`
}
