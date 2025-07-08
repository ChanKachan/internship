package products

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"net/http"
)

func (p *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	dbpool := database.NewProductDB(p.dbpool)

	product, err := dbpool.GetProducts()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(product)
}
