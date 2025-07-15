package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"internship/internal/models"
)

type ScanRows interface {
	ScanRowsOfWarehouse() error
	ScanRowsOfInventory() error
}

type scanRows struct {
	rows pgx.Rows
}

func NewScanRows(rows pgx.Rows) *scanRows {
	return &scanRows{rows}
}

func (s *scanRows) ScanRowsOfWarehouse() ([]models.Warehouse, error) {
	warehouses := []models.Warehouse{}
	var id, addressID uuid.UUID
	var city, street, building string

	for s.rows.Next() {
		s.rows.Scan(&id, &addressID, &city, &street, &building)
		address := models.Address{addressID, city, street, building}
		warehouse := models.Warehouse{id, address}
		warehouses = append(warehouses, warehouse)
	}

	if err := s.rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

func (s *scanRows) ScanRowsOfInventory() ([]models.ProductsBriefInfoInWarehouse, error) {
	inventory := []models.ProductsBriefInfoInWarehouse{}
	var productID uuid.UUID
	var productName string
	var price, discounted_price int

	for s.rows.Next() {
		s.rows.Scan(&productID, &productName, &price, &discounted_price)
		inventoryItems := models.ProductsBriefInfoInWarehouse{ProductId: productID, ProductName: productName, Price: price, PriceWithDiscount: discounted_price}
		inventory = append(inventory, inventoryItems)
	}

	if err := s.rows.Err(); err != nil {
		return nil, err
	}

	return inventory, nil
}
