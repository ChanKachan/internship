package models

import "github.com/google/uuid"

type ProductsBriefInfoInWarehouse struct {
	ProductId         uuid.UUID `json:"product_id"`
	ProductName       string    `json:"product_name"`
	Price             int       `json:"price"`
	PriceWithDiscount int       `json:"price_with_discount"`
}
