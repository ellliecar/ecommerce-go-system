```go
package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"ecommerce-go-system/internal/api"
	"ecommerce-go-system/internal/models"
	"ecommerce-go-system/internal/services"
)

// setupTestServer crea un servidor de prueba con dependencias inyectadas
func setupTestServer(t *testing.T) (*api.Server, *httptest.Server) {
	t.Helper()
	svc := services.NewOrderService()

	// Registrar productos de prueba
	p1, err := models.NewProduct("P001", "Laptop Pro", 1200.0, 50)
	if err != nil {
		t.Fatalf("failed to create test product: %v", err)
	}
	p2, err := models.NewProduct("P002", "Mouse Ergo", 45.0, 200)
	if err != nil {
		t.Fatalf("failed to create test product: %v", err)
	}
	svc.AddProduct(p1)
	svc.AddProduct(p2)

	server := api.NewServer(svc)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Enrutador minimalista para pruebas
		switch {
		case r.URL.Path == "/health" && r.Method == http.MethodGet:
			server.HealthCheck(w, r)
		case r.URL.Path == "/api/products" && r.Method == http.MethodGet:
			server.ListProducts(w, r)
		case strings.HasPrefix(r.URL.Path, "/api/products/") && r.Method == http.MethodGet:
			server.GetProduct(w, r)
		case r.URL.Path == "/api/orders" && r.Method == http.MethodPost:
			server.CreateOrder(w, r)
		case r.URL.Path == "/api/payments" && r.Method == http.MethodPost:
			server.ProcessPayment(w, r)
		case r.URL.Path == "/api/users/history" && r.Method == http.MethodGet:
			server.GetUserHistory(w, r)
		case r.URL.Path == "/api/inventory" && r.Method == http.MethodPut:
			server.UpdateInventory(w, r)
		case r.URL.Path == "/api/analytics" && r.Method == http.MethodGet:
			server.GetAnalytics(w, r)
		case r.URL.Path == "/api/concurrent" && r.Method == http.MethodPost:
			server.HandleConcurrency(w, r)
		default:
			http.NotFound(w, r)
		}
	}))

	return server, testServer
}

// TestHealthCheck valida el endpoint de estado del servidor
func TestHealthCheck(t *testing.T) {
	_, ts := setupTestServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if body["status"] != "healthy" {
		t.Errorf("expected status 'healthy', got '%v'", body["status"])
	}
	if _, ok := body["requests_processed"]; !ok {
		t.Error("expected 'requests_processed' field in response")
	}
}

// TestCreateOrder valida el endpoint de creación de pedidos
func TestCreateOrder(t *testing.T) {
	_, ts := setupTestServer(t)
	defer ts.Close()

	// Caso: pedido válido
	payload := map[string]interface{}{
		"product_id": "P001",
		"quantity":   2,
		"user_id":    "U100",
	}
	bodyBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/orders", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201, got %d", resp.StatusCode)
	}

	var respBody map[string]string
	json.NewDecoder(resp.Body).Decode(&respBody)
	if !strings.Contains(respBody["message"], "Order registered") {
		t.Errorf("expected success message, got: %v", respBody)
	}
}

// TestCreateOrder_InvalidProduct valida error de producto no encontrado
func TestCreateOrder_InvalidProduct(t *testing.T) {
	_, ts := setupTestServer(t)
	defer ts.Close()

	payload := map[string]interface{}{
		"product_id": "P999",
		"quantity":   1,
		"user_id":    "U100",
	}
	bodyBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/orders", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", resp.StatusCode)
	}
}

// TestCreateOrder_InsufficientStock valida error de stock insuficiente
func TestCreateOrder_InsufficientStock(t *testing.T) {
	_, ts := setupTestServer(t)
	defer ts.Close()

	payload := map[string]interface{}{
		"product_id": "P001",
		"quantity":   999,
		"user_id":    "U100",
	}
	bodyBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/orders", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400 for insufficient stock, got %d", resp.StatusCode)
	}
}

// TestConcurrency_Demo valida el endpoint de concurrencia
func TestConcurrency_Demo(t *testing.T) {
	_, ts := setupTestServer(t)
	defer ts.Close()

	payload := map[string]int{"concurrent_requests": 50}
	bodyBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/concurrent", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("concurrency request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode concurrency response: %v", err)
	}

	// Validaciones básicas de métricas
	processed, ok := result["processed"].(float64)
	if !ok {
		t.Error("expected 'processed' as number in response")
	}
	if processed < 45 {
		t.Errorf("expected at least 45 processed requests, got %.0f", processed)
	}

	timeMs, ok := result["time_ms"].(float64)
	if !ok {
		t.Error("expected 'time_ms' as number in response")
	}
	if timeMs > 100 {
		t.Errorf("expected response time <100ms, got %.0fms", timeMs)
	}

	t.Logf("Concurrencia: %.0f reqs procesados en %.0fms", processed, timeMs)
}

// TestJSONSerialization valida que todas las respuestas sean JSON válidos
func TestJSONSerialization(t *testing.T) {
	_, ts := setupTestServer(t)
	defer ts.Close()

	endpoints := []struct {
		method string
		path   string
		body   []byte
	}{
		{"GET", "/health", nil},
		{"GET", "/api/products", nil},
		{"GET", "/api/analytics", nil},
		{"POST", "/api/orders", []byte(`{"product_id":"P001","quantity":1,"user_id":"U1"}`)},
	}

	for _, ep := range endpoints {
		t.Run(ep.path, func(t *testing.T) {
			var req *http.Request
			var err error
			if ep.body != nil {
				req, err = http.NewRequest(ep.method, ts.URL+ep.path, bytes.NewReader(ep.body))
				req.Header.Set("Content-Type", "application/json")
			} else {
				req, err = http.NewRequest(ep.method, ts.URL+ep.path, nil)
			}
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer resp.Body.Close()

			contentType := resp.Header.Get("Content-Type")
			if !strings.Contains(contentType, "application/json") {
				t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
			}

			// Validar que el cuerpo es JSON parseable
			var raw json.RawMessage
			if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
				t.Errorf("response body is not valid JSON: %v", err)
			}
		})
	}
}
