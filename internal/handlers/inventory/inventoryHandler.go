package inventory

import "github.com/jackc/pgx/v5/pgxpool"

type inventoryHandler struct {
	dbpool *pgxpool.Pool
}

func NewInventoryHandler(dbpool *pgxpool.Pool) *inventoryHandler {
	return &inventoryHandler{dbpool: dbpool}
}
