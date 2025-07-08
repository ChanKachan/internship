package warehouses

import "github.com/jackc/pgx/v5/pgxpool"

type warehouseHandler struct {
	dbpool *pgxpool.Pool
}

func NewWarehouseHandler(dbpool *pgxpool.Pool) *warehouseHandler {
	return &warehouseHandler{dbpool: dbpool}
}
