package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"internship/internal/barcode"
	"internship/internal/logg"
	"internship/internal/models"
	"internship/internal/myuuid"
	"time"
)

type Product interface {
	CreateProduct(models.Product) error
	GetProducts() ([]models.Product, error)
	UpdateProduct(product Product) error
}

type productDB struct {
	dbpool *pgxpool.Pool
}

func NewProductDB(dbpool *pgxpool.Pool) *productDB {
	return &productDB{dbpool: dbpool}
}

func (p *productDB) CreateProduct(product models.Product) (models.Product, error) {
	logg.Logger.Info("Отправляю запрос на создание продукта в базу данных.",
		zap.String("package", "database.CreateProduct"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	product.ID = myuuid.GenerateUuid()

	defer cancel()

	_, err := p.dbpool.Exec(ctx,
		`INSERT INTO products VALUES ($1, $2, $3, $4, $5, $6)`,
		product.ID,
		product.Name,
		product.Description,
		product.Characteristic,
		product.Weight,
		barcode.Generate())

	if err != nil {
		return product, err
	}

	logg.Logger.Info("Запрос успешно завершен.",
		zap.String("package", "database.CreateProduct"))

	return product, nil
}

func (p *productDB) GetProducts() ([]models.Product, error) {
	logg.Logger.Info("Отправляю запрос на список продуктов в базу данных.",
		zap.String("package", "database.GetProducts"))

	logg.Logger.Debug("Отправляю SELECT запрос в таблицу products",
		zap.String("package", "database.GetProducts"))

	rows, err := p.dbpool.Query(context.Background(),
		`SELECT id,product_name,description,characteristics,weight,barcode FROM products`,
	)

	logg.Logger.Debug("SELECT запрос на таблицу products успешно завершен",
		zap.String("package", "database.GetProducts"))

	defer rows.Close()

	if err != nil {
		logg.Logger.Error(err.Error())
		return nil, err
	}

	products := []models.Product{}
	var characteristics json.RawMessage
	var id uuid.UUID
	var weight int
	var product_name, description, barcode string

	logg.Logger.Debug("Записываю данные в переменные.",
		zap.String("package", "database.GetProducts"))

	for rows.Next() {
		rows.Scan(&id, &product_name, &description, &characteristics, &weight, &barcode)
		product := models.Product{id, product_name, description, characteristics, weight, barcode}
		products = append(products, product)
	}

	logg.Logger.Debug("Все данные успешно записаны в переменную products.",
		zap.String("package", "database.GetProducts"))

	logg.Logger.Info("Запрос успешно завершен.",
		zap.String("package", "database.GetProducts"))

	return products, nil
}

// Метод получает данные о продукте. Данные на характеристику и описание проходят проверку на наличие,
// если все успешно, то мы обновляем данные о продукте.
func (p *productDB) UpdateProduct(product models.Product) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	logg.Logger.Info("Отправляю запрос на обновление данных продуктов в базу данных.",
		zap.String("package", "database.UpdateProduct"))

	logg.Logger.Debug("Проверка на наличие данных в переданной структуре.",
		zap.String("package", "database.UpdateProduct"))

	if product.Description == "" || product.Characteristic == nil || product.ID == uuid.Nil {
		err = fmt.Errorf("Характеристика и описания не должны быть пустыми!")

		return err
	}

	logg.Logger.Debug("Проверка на наличие данных в переданной структуре пройдена.",
		zap.String("package", "database.UpdateProduct"))

	logg.Logger.Debug("Отправляю запрос на обновление данных в базу данных.",
		zap.String("package", "database.UpdateProduct"))
	_, err = p.dbpool.Exec(ctx,
		`UPDATE products SET description = $1, characteristics = $2::jsonb WHERE id = $3`,
		product.Description,
		product.Characteristic,
		product.ID)

	if err != nil {
		return err
	}

	logg.Logger.Debug("Запрос успешно завершен.",
		zap.String("package", "database.UpdateProduct"))

	logg.Logger.Info("Данные обновлены.",
		zap.String("package", "database.UpdateProduct"))

	return nil
}
