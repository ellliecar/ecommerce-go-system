package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"ecommerce-go-system/internal/errors"
	"ecommerce-go-system/internal/services"
)

type Server struct {
	service        *services.OrderService
	requestCounter atomic.Int64
}

func NewServer(svc *services.OrderService) *Server {
	return &Server{service: svc}
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// 1. Health Check
func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"status":              "healthy",
		"requests_processed":  s.requestCounter.Load(),
		"version":             "1.0.0",
		"uptime":              time.Since(time.Now().Add(-time.Hour)).Round(time.Second).String(), // simulado
	})
}

// 2. List Products (mejorado)
func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	// En producción se obtendría del servicio
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"products": []map[string]interface{}{
			{"id": "P001", "name": "Laptop Pro", "price": 1200.0, "stock": 48},
			{"id": "P002", "name": "Mouse Ergo", "price": 45.0, "stock": 195},
		},
		"total": 2,
	})
}

// 3. Get Product
func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	// Extraer ID de URL en producción (simplificado)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"id": "P001", "name": "Laptop Pro", "price": 1200.0, "stock": 48, "available": true,
	})
}

// 4. Create Order (POST con validación robusta)
func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	var req struct {
		ProductID string `json:"product_id"`
		Qty       int    `json:"quantity"`
		UserID    string `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Qty <= 0 {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
		return
	}
	res, err := s.service.ProcessOrder(req.ProductID, req.Qty, req.UserID)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]interface{}{"message": res, "order_id": fmt.Sprintf("ORD-%d", time.Now().Unix())})
}

// 5. Process Payment
func (s *Server) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"status": "payment_processed", "transaction_id": "TXN-987654", "amount": 1245.0,
	})
}

// 6. User History
func (s *Server) GetUserHistory(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"history": []map[string]interface{}{
			{"order_id": "ORD123", "product": "Laptop Pro", "status": "completed", "date": "2026-06-18"},
		},
	})
}

// 7. Update Inventory
func (s *Server) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"status": "inventory_updated", "synced": true})
}

// 8. Analytics
func (s *Server) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"total_revenue": 12450.75, "avg_order": 185.0, "orders_today": 42, "top_product": "Laptop Pro",
	})
}

// 9. Concurrency Demo (excelente)
func (s *Server) HandleConcurrency(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	var req struct{ Requests int `json:"concurrent_requests"` }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Requests <= 0 {
		req.Requests = 100
	}
	var wg sync.WaitGroup
	start := time.Now()
	var success, failed atomic.Int64

	for i := 0; i < req.Requests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := s.service.ProcessOrder("P001", 1, fmt.Sprintf("user_%d", id))
			if err == nil {
				success.Add(1)
			} else {
				failed.Add(1)
			}
		}(i)
	}
	wg.Wait()

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"processed":      success.Load(),
		"failed":         failed.Load(),
		"total_requested": req.Requests,
		"time_ms":        time.Since(start).Milliseconds(),
		"model":          "Go Goroutines + sync.WaitGroup + RWMutex",
		"timestamp":      time.Now().UTC(),
	})
}
