package models

import "github.com/google/uuid"

type Analytic struct {
	WarehouseId     uuid.UUID `json:"warehouse_id"`
	ProductId       string    `json:"product_id"`
	QuantityOfSales int       `json:"quantity_of_products_sold"`
	TotalPrice      int       `json:"price_of_products_sold"`
}
