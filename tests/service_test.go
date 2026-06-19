package tests

import (
	"testing"
	"ecommerce-go-system/internal/models"
	"ecommerce-go-system/internal/services"
)

func TestOrderService_ProcessOrder(t *testing.T) {
	svc := services.NewOrderService()
	p, _ := models.NewProduct("P1", "Item", 10.0, 5)
	svc.AddProduct(p)

	// Caso éxito
	msg, err := svc.ProcessOrder("P1", 2, "U1")
	if err != nil || msg == "" {
		t.Fatalf("expected success, got err: %v", err)
	}

	// Caso stock insuficiente
	_, err = svc.ProcessOrder("P1", 10, "U2")
	if err == nil {
		t.Fatal("expected insufficient stock error")
	}
}
