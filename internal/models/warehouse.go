package models

import "github.com/google/uuid"

type Warehouse struct {
	ID      uuid.UUID `json:"id"`
	Address Address   `json:"address"`
}
