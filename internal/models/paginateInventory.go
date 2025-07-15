package models

type PaginationProductsOfWarehouse struct {
	InfoBriefInventory []ProductsBriefInfoInWarehouse `json:"data"`
	Pagination         Pagination                     `json:"meta"`
}
