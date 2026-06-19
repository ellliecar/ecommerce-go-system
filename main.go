package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-go-system/internal/api"
	"ecommerce-go-system/internal/models"
	"ecommerce-go-system/internal/services"
)

func main() {
	svc := services.NewOrderService()

	// Registro inicial de productos
	p1, _ := models.NewProduct("P001", "Laptop Pro", 1200.0, 50)
	p2, _ := models.NewProduct("P002", "Mouse Ergo", 45.0, 200)
	svc.AddProduct(p1)
	svc.AddProduct(p2)

	server := api.NewServer(svc)

	// Rutas (8 servicios + 1 endpoint de concurrencia)
	http.HandleFunc("/health", server.HealthCheck)
	http.HandleFunc("/api/products", server.ListProducts)
	http.HandleFunc("/api/products/", server.GetProduct)
	http.HandleFunc("/api/orders", server.CreateOrder)
	http.HandleFunc("/api/payments", server.ProcessPayment)
	http.HandleFunc("/api/users/history", server.GetUserHistory)
	http.HandleFunc("/api/inventory", server.UpdateInventory)
	http.HandleFunc("/api/analytics", server.GetAnalytics)
	http.HandleFunc("/api/concurrent", server.HandleConcurrency)

	fmt.Println("🚀 Servidor Go activo en http://localhost:8080")
	fmt.Println("📦 8 Servicios Web + Concurrencia nativa + JSON serializado")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
