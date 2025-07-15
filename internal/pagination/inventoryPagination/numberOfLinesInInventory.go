package inventoryPagination

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"internship/internal/logg"
)

type PaginationCounter interface {
	InventoryCount() int
}

type paginationItems struct {
	dbpool *pgxpool.Pool
}

func NewPaginationCounter(dbpool *pgxpool.Pool) *paginationItems {
	return &paginationItems{dbpool: dbpool}
}

func (p *paginationItems) inventoryCount(warehouseID uuid.UUID) int {
	var count int
	logg.Logger.Info("Запрос на количество кортежей в таблицу inventory.",
		zap.String("package", "inventoryPagination.inventoryCount"))

	err := p.dbpool.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM inventory WHERE warehouse_id = $1`, warehouseID).Scan(&count)

	if err != nil {
		logg.Logger.Error(err.Error(),
			zap.String("package", "inventoryPagination.inventoryCount"))
		return 0
	}

	logg.Logger.Info("Запрос успешно завершен.",
		zap.String("package", "inventoryPagination.inventoryCount"))
	return count
}
