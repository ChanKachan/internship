package main

import (
	"fmt"
	"internship/internal/database"
	"internship/internal/handlers"
	"internship/internal/handlers/products"
	"internship/internal/handlers/warehouses"
	"internship/internal/logg"
	"log"
	"net/http"
	"os"
)

func main() {
	logg.InitLogger()
	dbpool := database.ConnectDatabase()

	defer dbpool.Close()
	defer logg.Logger.Sync()

	dir, _ := os.Getwd()
	log.Println("Текущая директория:", dir)

	http.HandleFunc("/api/heath", handlers.Heath)
	//http.Handle("/api/create_warehouse", middleware.CreateAddressForWarehouses(http.HandlerFunc(warehouses.CreateWarehouses)))
	warehouse := warehouses.NewWarehouseHandler(dbpool)
	http.HandleFunc("/api/create_warehouse", warehouse.CreateWarehouses)
	http.HandleFunc("/api/warehouses", warehouse.GetWarehouses)

	product := products.NewProductHandler(dbpool)
	http.HandleFunc("/api/create_product", product.CreateProduct)
	http.HandleFunc("/api/products", product.GetProducts)
	http.HandleFunc("/api/update_product", product.PutProductCharacteristicOrDescription)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
