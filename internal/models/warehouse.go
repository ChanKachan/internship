package models

type Warehouse struct {
	ID      string  `json:"id"`
	Address Address `json:"address"`
}
