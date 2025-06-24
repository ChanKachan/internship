package models

type Analytic struct {
	WarehouseId     string `json:"warehouse_id"`
	ProductId       string `json:"product_id"`
	QuantityOfSales int    `json:"quantity_of_products_sold"`
	TotalPrice      int    `json:"price_of_products_sold"`
}

type SDsda struct {
	wafds string `json:"wafds"`
}
