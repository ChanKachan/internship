package products

import "github.com/jackc/pgx/v5/pgxpool"

type productHandler struct {
	dbpool *pgxpool.Pool
}

func NewProductHandler(dbpool *pgxpool.Pool) *productHandler {
	return &productHandler{dbpool: dbpool}
}
