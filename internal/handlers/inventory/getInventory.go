package inventory

import (
	"encoding/json"
	"internship/internal/database"
	"internship/internal/models"
	"net/http"
)

// todo Нужно метод работает без пагинации.
// нужно добавить пагинацию
// todo структура выводит лишнюю инфу
// нужно решить чтобы выводились конкретные значения и переделать get для пагинации
func (i *inventoryHandler) GetInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//offset := r.URL.Query().Get("offset")
	//limit := r.URL.Query().Get("limit")

	dbpool := database.NewInventoryDB(i.dbpool)

	inventory, err := dbpool.GetProductsFromWarehouse(10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorItem{Message: err.Error(), Code: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(inventory)
}
