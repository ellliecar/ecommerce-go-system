package services

import (
	"fmt"
	"sync"
	"ecommerce-go-system/internal/errors"
	"ecommerce-go-system/internal/interfaces"
)

// OrderService: Orquestador con concurrencia segura (sync.RWMutex)
type OrderService struct {
	mu                sync.RWMutex
	catalog           map[string]interfaces.IProduct
	orderQueue        []string
	processedSessions map[string]bool
}

func NewOrderService() *OrderService {
	return &OrderService{
		catalog:           make(map[string]interfaces.IProduct),
		orderQueue:        make([]string, 0),
		processedSessions: make(map[string]bool),
	}
}

func (os *OrderService) AddProduct(p interfaces.IProduct) error {
	os.mu.Lock()
	defer os.mu.Unlock()
	if _, exists := os.catalog[p.GetID()]; exists {
		return fmt.Errorf("product %s already exists", p.GetID())
	}
	os.catalog[p.GetID()] = p
	return nil
}

// ProcessOrder: Validación → Mutación controlada → Encolado → Trazabilidad
func (os *OrderService) ProcessOrder(productID string, qty int, userID string) (string, error) {
	sessionKey := fmt.Sprintf("%s::%s", productID, userID)

	os.mu.RLock()
	if os.processedSessions[sessionKey] {
		os.mu.RUnlock()
		return "", errors.ErrDuplicateRequest
	}
	product, exists := os.catalog[productID]
	if !exists {
		os.mu.RUnlock()
		return "", errors.ErrProductNotFound
	}
	if !product.IsAvailable(qty) {
		os.mu.RUnlock()
		return "", errors.ErrInsufficientStock
	}
	os.mu.RUnlock()

	err := product.ReduceStock(qty)
	if err != nil {
		return "", fmt.Errorf("%w: %v", errors.ErrPaymentFailed, err)
	}

	os.mu.Lock()
	os.processedSessions[sessionKey] = true
	os.orderQueue = append(os.orderQueue, fmt.Sprintf("%s|%d|%s", productID, qty, userID))
	queueSize := len(os.orderQueue)
	os.mu.Unlock()

	return fmt.Sprintf("Order registered. Queue size: %d", queueSize), nil
}

func (os *OrderService) GetQueueStatus() int {
	os.mu.RLock()
	defer os.mu.RUnlock()
	return len(os.orderQueue)
}
