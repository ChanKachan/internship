package models

import "github.com/google/uuid"

type Address struct {
	ID       uuid.UUID `json:"id"`
	City     string    `json:"city"`
	Street   string    `json:"street"`
	Building string    `json:"building"`
}
