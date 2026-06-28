package tests

import (
	"testing"
)

func TestExample(t *testing.T) {
	// Prueba básica temporal
	if 1+1 != 2 {
		t.Error("Error básico")
	}
	t.Log("Prueba básica OK")
}
