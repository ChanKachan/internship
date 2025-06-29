package warehouses

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"internship/internal/uuid"
	"net/http"
	"strconv"
)

func CreateWarehouses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	warehouse := models.Warehouse{}
	dbpool := database.NewWarehouseDB(database.ConnectDatabase())

	err := json.NewDecoder(r.Body).Decode(&warehouse)

	warehouse.Address.ID = uuid.GenerateUuid()
	warehouse.ID = uuid.GenerateUuid()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	err = dbpool.CreateWarehouseWithAddress(warehouse)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{ErrorItem: err, Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"warehouse_id": warehouse.ID, "message": "Склад успешно добавлен!", "code": strconv.Itoa(http.StatusOK),
	})
}
