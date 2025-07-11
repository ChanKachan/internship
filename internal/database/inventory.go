package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"internship/internal/logg"
	"internship/internal/models"
	"time"
)

/*
0) Создание связи товара и склада: Указание цены на конкретном складе;
1) Обновление количества определённого товара на складе увеличение (постапление товара на склад);
2) Создание скидку на опеределённый список товаров на складе;
3) Получение списка товаров по конкретному складу с пагинацией в элементе списка должна содержаться информация: идентификатор, наименование товара, цена без скидки цена со скидкой;
4) Получения товара на складе: получения ответе все информации по товару на складе: идентификатор товара, наименование, описание, характеристики, штрих код, цена, цена со скидкой, количество;
5) Получения подсчёта: запрос который получает в себе индектификатор склада, идектификаторы товаров и их количество. Вернуть в ответе подсумировку цен списка.
6) Покупка товаров со склада: запрос который получает в себе индектификатор склада, идектификаторы товаров и их количество. Уменьшает количетсво товаров на складе. Если какой то товар был уже продан и количества не хватает вернуть ошибку.
*/

type PostInventory interface {
	CreateInventory(inventory models.Inventory) error
}

type PutInventory interface {
	UpdateQuantity(inventory models.Inventory) error
	UpdateDiscount(inventory models.Inventory) error
	ReduceStockItems(inventory []models.Inventory) error
}

type GetInventory interface {
	GetProductsFromWarehouse(inventory models.Inventory) (models.Inventory, error)
	GetProductInformationInStock(inventory models.Inventory) (models.Inventory, error)
	GetCostOfProductInStock(inventory models.Inventory) (models.Inventory, error)

type inventoryDB struct {
	dbpool *pgxpool.Pool
}

func NewInventoryDB(dbpool *pgxpool.Pool) *inventoryDB {
	return &inventoryDB{dbpool: dbpool}
}

func (i *inventoryDB) CreateInventory(inventory models.Inventory) error {
	logg.Logger.Info("Отправляю запрос на создание связи товара и склада в базу данных.",
		zap.String("package", "database.CreateInventory"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := i.dbpool.Exec(ctx,
		`INSERT INTO inventory (warehouse_id,product_id,price) VALUES ($1, $2, $3)`,
		inventory.WarehouseId,
		inventory.ProductId,
		inventory.Price)

	if err != nil {
		logg.Logger.Error(err.Error(), zap.String("package", "database.CreateInventory"))
		return err
	}

	logg.Logger.Info("Запрос успешно завершен.",
		zap.String("package", "database.CreateInventory"))

	return nil
}

func (i *inventoryDB) UpdateQuantity(inventory models.Inventory) error {
	logg.Logger.Info("Запрос на обновление количество товара на складе.",
		zap.String("package", "database.UpdateQuantity"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := i.dbpool.Exec(ctx,
		`UPDATE inventory SET quantity_of_product = $1 WHERE product_id = $2 AND warehouse_id = $3`,
		inventory.Quantity,
		inventory.ProductId,
		inventory.WarehouseId)

	if err != nil {
		logg.Logger.Error(err.Error(), zap.String("package", "database.UpdateQuantity"))
		return err
	}

	logg.Logger.Info("Запрос успешно завершен.",
		zap.String("package", "database.UpdateQuantity"))

	return nil
}

func (i *inventoryDB) UpdateDiscount(inventory models.Inventory) error {
	logg.Logger.Info("Запрос на скидку для товара в конкретном складе.",
		zap.String("package", "database.UpdateDiscount"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := i.dbpool.Exec(ctx,
		`UPDATE inventory SET percentage_discount_from_price = $1 WHERE product_id = $2 AND warehouse_id = $3`,
		inventory.Percentage,
		inventory.ProductId,
		inventory.WarehouseId)

	if err != nil {
		logg.Logger.Error(err.Error(), zap.String("package", "database.UpdateDiscount"))
		return err
	}

	logg.Logger.Info("Запрос успешно завершен.",
		zap.String("package", "database.UpdateDiscount"))

	return nil
}
