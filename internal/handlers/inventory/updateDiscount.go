package inventory

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"net/http"
	"strconv"
)

func (i *inventoryHandler) UpdateDiscount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Allow", "PUT")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var inventory models.Inventory

	dbpool := database.NewInventoryDB(i.dbpool)

	err := json.NewDecoder(r.Body).Decode(&inventory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	err = dbpool.UpdateDiscount(inventory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"product_id":   inventory.ProductId.String(),
		"warehouse_id": inventory.WarehouseId.String(),
		"message":      "Процент скидки на товар в складе изменен!",
		"code":         strconv.Itoa(http.StatusOK),
	})
}
