package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"internship/internal/logg"
	"internship/internal/models"
)

type Warehouse interface {
	CreateWarehouseWithAddress(warehouse models.Warehouse) error
	GetByID(id string) (Warehouse, error)
	Update(warehouse Warehouse) error
}

type warehouseDB struct {
	dbpool *pgxpool.Pool
}

func NewWarehouseDB(dbpool *pgxpool.Pool) *warehouseDB {
	return &warehouseDB{dbpool: dbpool}
}

func (w warehouseDB) CreateWarehouseWithAddress(warehouse models.Warehouse) error {
	logg.Logger.Info("Добавляем склад в бд")

	tx, err := w.dbpool.Begin(context.Background())
	defer w.dbpool.Close()
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`INSERT INTO address (id, city,street,building) VALUES ($1,$2,$3,$4)`,
		warehouse.Address.ID,
		warehouse.Address.City,
		warehouse.Address.Street,
		warehouse.Address.Building)

	if err != nil {
		logg.Logger.Error(err.Error())
		return err
	}

	_, err = tx.Exec(context.Background(),
		`INSERT INTO warehouses (id, address_id) VALUES ($1,$2)`,
		warehouse.ID,
		warehouse.Address.ID)

	err = tx.Commit(context.Background())
	if err != nil {
		logg.Logger.Error(err.Error())
		return err
	}

	logg.Logger.Info("Склад добавлен в бд!")
	return nil
}
