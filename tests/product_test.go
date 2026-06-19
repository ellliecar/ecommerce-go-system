package tests

import (
	"testing"
	"ecommerce-go-system/internal/models"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		name    string
		price   float64
		stock   int
		wantErr bool
	}{
		{"Valid", 100.0, 10, false},
		{"InvalidPrice", -5.0, 10, true},
		{"InvalidStock", 50.0, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := models.NewProduct("T1", "Test", tt.price, tt.stock)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
