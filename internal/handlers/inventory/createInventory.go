package inventory

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"net/http"
	"strconv"
)

func CreateInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var inventory models.Inventory

	dbpool := database.NewInventoryDB(database.ConnectDatabase())

	err := json.NewDecoder(r.Body).Decode(&inventory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	err = dbpool.CreateInventory(inventory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"product_id":   inventory.ProductId.String(),
		"warehouse_id": inventory.WarehouseId.String(),
		"message":      "Cвязь товара и склада успешно добавлен!",
		"code":         strconv.Itoa(http.StatusOK),
	})
}
