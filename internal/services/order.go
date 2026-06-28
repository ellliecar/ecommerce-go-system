package services

import (
	"fmt"
	"sync"

	"ecommerce-go-system/internal/models"
)

type OrderService struct {
	products map[string]*models.Product
	mu       sync.RWMutex
}

func NewOrderService() *OrderService {
	return &OrderService{
		products: make(map[string]*models.Product),
	}
}

func (s *OrderService) AddProduct(p *models.Product) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.products[p.GetID()] = p
}

func (s *OrderService) ProcessOrder(productID string, qty int, userID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, exists := s.products[productID]
	if !exists {
		return "", fmt.Errorf("producto no encontrado")
	}

	if err := product.ReduceStock(qty); err != nil {
		return "", err
	}

	return fmt.Sprintf("Pedido procesado para usuario %s", userID), nil
}
