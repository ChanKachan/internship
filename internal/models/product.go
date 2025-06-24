package models

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"product_name"`
	Description string `json:"description"`
	//characteristic в бд создан как json. Нужно его реализовать
	Weight  int    `json:"weight"`
	Barcode string `json:"barcode"`
}
