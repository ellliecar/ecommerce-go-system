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

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"status":             "healthy",
		"requests_processed": s.requestCounter.Load(),
	})
}

func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"products": []string{"P001", "P002"}})
}

func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"id": "P001", "name": "Laptop Pro"})
}

func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	var req struct {
		ProductID string `json:"product_id"`
		Qty       int    `json:"quantity"`
		UserID    string `json:"user_id"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	res, err := s.service.ProcessOrder(req.ProductID, req.Qty, req.UserID)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	jsonResponse(w, http.StatusCreated, map[string]interface{}{"message": res})
}

func (s *Server) HandleConcurrency(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	var wg sync.WaitGroup
	var success atomic.Int64
	start := time.Now()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.service.ProcessOrder("P001", 1, "user")
			if err == nil {
				success.Add(1)
			}
		}()
	}
	wg.Wait()

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"processed": success.Load(),
		"time_ms":   time.Since(start).Milliseconds(),
	})
}
