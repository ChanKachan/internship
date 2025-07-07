package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"internship/internal/models"
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

type Inventory interface {
	CreateInventory(inventory models.Inventory) error
	UpdateQuantity(inventory models.Inventory) error
	UpdateDiscount(inventory models.Inventory) error
	GetProductsFromWarehouse(inventory models.Inventory) (models.Inventory, error)
	GetProductInformationInStock(inventory models.Inventory) (models.Inventory, error)
	GetCostOfProductInStock(warehouseID, productID uuid.UUID, count int) (models.Inventory, error)
	ReduceStockItems(inventory []models.Inventory) error
}

type inventoryDB struct {
	dbpool *pgxpool.Pool
}

func NewInventoryDB(dbpool *pgxpool.Pool) *inventoryDB {
	return &inventoryDB{dbpool: dbpool}
}

func (i *inventoryDB) CreateInventory(inventory models.Inventory) error {
	return nil
}
