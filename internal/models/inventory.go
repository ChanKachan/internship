package models

type Inventory struct {
	WarehouseId string `json:"warehouse_id"`
	ProductId   string `json:"product_id"`
	Quantity    int    `json:"quantity_of_product"`
	Price       int    `json:"price"`
	Percentage  int    `json:"percentage_discount_from_price"`
}
