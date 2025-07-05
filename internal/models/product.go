package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Product struct {
	ID             uuid.UUID       `json:"id"`
	Name           string          `json:"product_name"`
	Description    string          `json:"description"`
	Characteristic json.RawMessage `json:"characteristic"`
	Weight         int             `json:"weight"`
	Barcode        string          `json:"barcode"`
}
