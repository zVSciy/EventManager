package main

import (
	"fmt"
	"io"
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

func TestGetPayment(t *testing.T) {
	var tests = []struct {
		id           string
		wantCode     int
		wantResponse string
	}{
		{"this-id-is-invalid", http.StatusBadRequest, `{"error":"INVALID_PAYMENT_ID"}`},
		{"hallo", http.StatusBadRequest, `{"error":"INVALID_PAYMENT_ID"}`},
		{"6749909e35593ccce69c82a9", http.StatusInternalServerError, `{"error":"DATABASE_NOT_INITIALIZED"}`},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("GET /payments/{id}", handlers.GetPayment)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	for _, tt := range tests {

		testname := tt.id
		t.Run(testname, func(t *testing.T) {
			path := fmt.Sprintf("/payments/%v", tt.id)

			res, err := http.Get(ts.URL + path)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			if status := res.StatusCode; status != tt.wantCode {
				t.Errorf("got %v want %v", status, tt.wantCode)
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if body := strings.TrimSpace(string(resBody)); body != tt.wantResponse {
				t.Errorf("got %v want %v", body, tt.wantResponse)
			}
		})
	}
}

func TestFallback(t *testing.T) {
	var tests = []struct {
		name         string
		wantCode     int
		wantResponse string
	}{
		{"Fallback", http.StatusNotFound, `{"error":"Not Found"}`},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("/", handlers.NotFound)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			path := "/"
			res, err := http.Get(ts.URL + path)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			if status := res.StatusCode; status != tt.wantCode {
				t.Errorf("got %v want %v", status, tt.wantCode)
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			if body := strings.TrimSpace(string(resBody)); body != tt.wantResponse {
				t.Errorf("got %v want %v", body, tt.wantResponse)
			}
		})
	}
}
