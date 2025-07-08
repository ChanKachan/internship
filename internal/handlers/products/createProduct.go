package products

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"net/http"
	"strconv"
)

func (p *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var product models.Product

	dbpool := database.NewProductDB(p.dbpool)

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	product, err = dbpool.CreateProduct(product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"product_id": product.ID.String(), "message": "Продукт успешно добавлен!", "code": strconv.Itoa(http.StatusOK),
	})
}
