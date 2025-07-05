package main

import (
	"fmt"
	"internship/internal/handlers"
	"internship/internal/handlers/warehouses"
	"internship/internal/logg"
	"log"
	"net/http"
	"os"
)

func main() {
	logg.InitLogger()

	defer logg.Logger.Sync()

	dir, _ := os.Getwd()
	log.Println("Текущая директория:", dir)

	http.HandleFunc("/api/heath", handlers.Heath)
	//http.Handle("/api/create_warehouse", middleware.CreateAddressForWarehouses(http.HandlerFunc(warehouses.CreateWarehouses)))
	http.HandleFunc("/api/create_warehouse", warehouses.CreateWarehouses)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
