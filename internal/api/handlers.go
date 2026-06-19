package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"ecommerce-go-system/internal/services"
)

type Server struct {
	service        *services.OrderService
	requestCounter atomic.Int64
}

func NewServer(svc *services.OrderService) *Server {
	return &Server{service: svc}
}

// jsonResponse helper para serialización JSON explícita
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// 1. Health Check
func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"status":           "healthy",
		"requests_processed": s.requestCounter.Load(),
	})
}

// 2. List Products
func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"message": "Catalog ready", "format": "JSON"})
}

// 3. Get Product
func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"id": "P001", "name": "Laptop Pro", "price": 1200.0})
}

// 4. Create Order
func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	var req struct {
		ProductID string `json:"product_id"`
		Qty       int    `json:"quantity"`
		UserID    string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
		return
	}
	res, err := s.service.ProcessOrder(req.ProductID, req.Qty, req.UserID)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]string{"message": res})
}

// 5. Process Payment
func (s *Server) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]string{"status": "payment_processed", "serializaed": "JSON"})
}

// 6. User History
func (s *Server) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, []map[string]string{{"order_id": "ORD123", "status": "completed"}})
}

// 7. Update Inventory
func (s *Server) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]string{"status": "inventory_updated_synced"})
}

// 8. Analytics
func (s *Server) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"total_revenue": 4500.0, "avg_order": 120.0})
}

// 9. Concurrency Demo (Goroutines + WaitGroup + Mutex)
func (s *Server) HandleConcurrency(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	var req struct{ Requests int `json:"concurrent_requests"` }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Requests <= 0 {
		req.Requests = 50
	}

	var wg sync.WaitGroup
	start := time.Now()
	var success atomic.Int64

	for i := 0; i < req.Requests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := s.service.ProcessOrder("P001", 1, fmt.Sprintf("user_%d", id))
			if err == nil || err.Error() == errors.ErrInsufficientStock.Error() {
				success.Add(1)
			}
		}(i)
	}
	wg.Wait()

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"processed":       success.Load(),
		"total_requested": req.Requests,
		"time_ms":         time.Since(start).Milliseconds(),
		"model":           "Go Goroutines + sync.WaitGroup + sync.RWMutex",
		"timestamp":       time.Now().UTC(),
	})
}
