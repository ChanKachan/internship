package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"internship/internal/logg"
	"internship/internal/models"
)

type PostWarehouse interface {
	CreateWarehouseWithAddress(warehouse models.Warehouse) error
}

type GetWarehouse interface {
	GetWarehouses() ([]models.Warehouse, error)
}

type warehouseDB struct {
	dbpool *pgxpool.Pool
}

func NewWarehouseDB(dbpool *pgxpool.Pool) *warehouseDB {
	return &warehouseDB{dbpool: dbpool}
}

func (w warehouseDB) CreateWarehouseWithAddress(warehouse models.Warehouse) error {
	logg.Logger.Info("Отправляю запрос в базу данных.",
		zap.String("package", "database.CreateWarehouseWithAddress"))

	logg.Logger.Debug("Отправляю запрос в таблицу warehouses, транзакция запущена",
		zap.String("package", "database.CreateWarehouseWithAddress"))

	tx, err := w.dbpool.Begin(context.Background())

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

	logg.Logger.Debug("Транзакция завершина в таблице warehouses",
		zap.String("package", "database.CreateWarehouseWithAddress"))

	logg.Logger.Info("Склад добавлен в базу данных.",
		zap.String("package", "database.CreateWarehouseWithAddress"))
	return nil
}

func (w warehouseDB) GetWarehouses() ([]models.Warehouse, error) {
	logg.Logger.Info("Отправляю запрос SELECT в базу данных.",
		zap.String("package", "database.CreateWarehouseWithAddress"))

	rows, err := w.dbpool.Query(context.Background(),
		"SELECT w.id,w.address_id,a.city,a.street,a.building FROM warehouses w INNER JOIN address a ON a.id = w.address_id",
	)

	defer rows.Close()

	if err != nil {
		logg.Logger.Error(err.Error())
		return nil, err
	}

	warehouses := []models.Warehouse{}
	var id, addressID uuid.UUID
	var city, street, building string

	for rows.Next() {
		rows.Scan(&id, &addressID, &city, &street, &building)
		address := models.Address{addressID, city, street, building}
		warehouse := models.Warehouse{id, address}
		warehouses = append(warehouses, warehouse)
	}

	logg.Logger.Info("Запрос на поиск данных в таблице warehouses прошел успешно.",
		zap.String("package", "database.CreateWarehouseWithAddress"))

	return warehouses, nil
}
