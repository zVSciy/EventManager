package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zVSciy/EventManager/Payment/internal/handlers"
)

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.HealthCheck)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"healthy"}`
	if body := strings.TrimSpace(w.Body.String()); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetPayments(t *testing.T) {
	req := httptest.NewRequest("GET", "/accounts/testuser/payments", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.GetPayments)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expected := `{"error":"DATABASE_NOT_INITIALIZED"}`
	if body := strings.TrimSpace(w.Body.String()); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestFallback(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.NotFound)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	expected := `{"error":"Not Found"}`
	if body := strings.TrimSpace(w.Body.String()); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}
