package models

import "github.com/google/uuid"

type Inventory struct {
	WarehouseId       uuid.UUID `json:"warehouse_id"`
	ProductId         uuid.UUID `json:"product_id"`
	ProductName       string    `json:"product_name"`
	Quantity          int       `json:"quantity_of_product"`
	Price             int       `json:"price"`
	PriceWithDiscount int       `json:"price_with_discount"`
	Percentage        int       `json:"percentage_discount_from_price"`
}
