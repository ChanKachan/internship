package models

import "github.com/google/uuid"

type Inventory struct {
	WarehouseId uuid.UUID `json:"warehouse_id"`
	ProductId   uuid.UUID `json:"product_id"`
	Quantity    int       `json:"quantity_of_product"`
	Price       int       `json:"price"`
	Percentage  int       `json:"percentage_discount_from_price"`
}
