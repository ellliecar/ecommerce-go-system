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
		"status": "healthy",
	})
}

func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"message": "Lista de productos"})
}

func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{"id": "P001"})
}

func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusCreated, map[string]interface{}{"message": "Pedido creado"})
}

func (s *Server) HandleConcurrency(w http.ResponseWriter, r *http.Request) {
	s.requestCounter.Add(1)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Concurrencia ejecutada",
		"processed": 30,
	})
}
