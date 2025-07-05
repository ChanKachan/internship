package models

type Address struct {
	ID       string `json:"id"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Building string `json:"building"`
}
