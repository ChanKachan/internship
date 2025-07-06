package products

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"net/http"
	"strconv"
)

func PutProductCharacteristicOrDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", "PUT")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	dbpool := database.NewProductDB(database.ConnectDatabase())
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	err = dbpool.UpdateProduct(product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"product_id": product.ID.String(), "message": "Данные успешно изменены!", "code": strconv.Itoa(http.StatusOK),
	})
}
