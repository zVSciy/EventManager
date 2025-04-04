package main

import (
	"encoding/json"
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

func TestGetAccount(t *testing.T) {
	var tests = []struct {
		userID       string
		authHeader   string
		wantCode     int
		wantResponse string
	}{
		{
			"blabla",
			"Bearer eyJhbGciOi",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"abcabcabc",
			"Bearer NiIsInR5cC",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"hansi",
			"",
			http.StatusUnauthorized,
			`{"error":"Unauthorized"}`,
		},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("GET /accounts/{user_id}", handlers.GetAccount)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	client := &http.Client{}

	for _, tt := range tests {

		testname := tt.userID
		t.Run(testname, func(t *testing.T) {
			path := fmt.Sprintf("%v/accounts/%v", ts.URL, tt.userID)

			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", tt.authHeader)

			res, err := client.Do(req)
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

func TestGetPayments(t *testing.T) {
	var tests = []struct {
		userID       string
		authHeader   string
		wantCode     int
		wantResponse string
	}{
		{
			"testuser",
			"Bearer aniu9dofs",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"hallo",
			"Bearer 9efh8zirwu",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"lachs",
			"",
			http.StatusUnauthorized,
			`{"error":"Unauthorized"}`,
		},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("GET /accounts/{user_id}/payments", handlers.GetPayments)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	client := &http.Client{}

	for _, tt := range tests {

		testname := tt.userID
		t.Run(testname, func(t *testing.T) {
			path := fmt.Sprintf("%v/accounts/%v/payments", ts.URL, tt.userID)

			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", tt.authHeader)

			res, err := client.Do(req)
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

func TestGetPayment(t *testing.T) {
	var tests = []struct {
		id           string
		userID       string
		authHeader   string
		wantCode     int
		wantResponse string
	}{
		{
			"this-id-is-invalid",
			"fla",
			"Bearer a9ßh8gufq",
			http.StatusBadRequest,
			`{"error":"INVALID_PAYMENT_ID"}`,
		},
		{
			"hallo",
			"alf",
			"Bearer gre9h8ij",
			http.StatusBadRequest,
			`{"error":"INVALID_PAYMENT_ID"}`,
		},
		{
			"6749909e35593ccce69c82a9",
			"ali",
			"Bearer fa9n9sddss",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"60c72b2f9e1d4e2d88f3f79e",
			"oli",
			"",
			http.StatusUnauthorized,
			`{"error":"Unauthorized"}`,
		},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("GET /payments/{id}", handlers.GetPayment)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	client := &http.Client{}

	for _, tt := range tests {

		testname := tt.id
		t.Run(testname, func(t *testing.T) {
			path := fmt.Sprintf("%v/payments/%v", ts.URL, tt.id)

			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", tt.authHeader)

			res, err := client.Do(req)
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

func TestCreateAccount(t *testing.T) {
	var tests = []struct {
		userID       string
		authHeader   string
		wantCode     int
		wantResponse string
	}{
		{
			"meowmeow",
			"Bearer 9piubanodfs",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"lalalalala",
			"Bearer sdpoafs0sas",
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
		{
			"heyhoo",
			"",
			http.StatusUnauthorized,
			`{"error":"Unauthorized"}`,
		},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("POST /accounts", handlers.CreateAccount)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	client := &http.Client{}

	for _, tt := range tests {

		testname := tt.userID
		t.Run(testname, func(t *testing.T) {
			path := fmt.Sprintf("%v/accounts", ts.URL)

			payload := fmt.Sprintf(`{"user_id": "%v"}`, tt.userID)

			req, err := http.NewRequest("POST", path, strings.NewReader(payload))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Authorization", tt.authHeader)

			res, err := client.Do(req)
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

func TestCreatePayment(t *testing.T) {
	var tests = []struct {
		name           string
		idempotencyKey string
		authHeader     string
		payload        string
		wantCode       int
		wantResponse   string
	}{
		{
			"Missing Idempotency Key",
			"",
			"Bearer asdbp9fuio",
			`{"currency": "EUR", "amount": 100}`,
			http.StatusBadRequest,
			`{"error":"MISSING_IDEMPOTENCY_KEY"}`,
		},
		{
			"Missing Auth Header",
			"valid-idempotency-key",
			"",
			`{"currency": "EUR", "amount": 100}`,
			http.StatusUnauthorized,
			`{"error":"Unauthorized"}`,
		},
		{
			"Invalid JSON Body",
			"valid-idempotency-key",
			"Bearer asd90ißfn",
			`{"currency": "EUR", "amount": "invalid"}`,
			http.StatusBadRequest,
			`{"error":"Bad Request"}`,
		},
		{
			"Bad Currency",
			"valid-idempotency-key",
			"Bearer a9isdfn",
			`{"currency": "USD", "amount": 100}`,
			http.StatusBadRequest,
			`{"error":"BAD_CURRENCY_USE_EUR"}`,
		},
		{
			"Invalid Amount",
			"valid-idempotency-key",
			"Bearer sa9dnf0",
			`{"currency": "EUR", "amount": -1}`,
			http.StatusBadRequest,
			`{"error":"INVALID_AMOUNT"}`,
		},
		{
			"Database Not Initialized",
			"valid-idempotency-key",
			"Bearer asd9ßf8b",
			`{"currency": "EUR", "amount": 100}`,
			http.StatusInternalServerError,
			`{"error":"DATABASE_NOT_INITIALIZED"}`,
		},
	}

	testMux := http.NewServeMux()
	testMux.HandleFunc("/payments", handlers.CreatePayment)
	ts := httptest.NewServer(testMux)
	defer ts.Close()

	client := &http.Client{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("%v/payments", ts.URL)

			req, err := http.NewRequest("POST", path, strings.NewReader(tt.payload))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Idempotency-Key", tt.idempotencyKey)
			req.Header.Set("Authorization", tt.authHeader)

			res, err := client.Do(req)
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

			var response map[string]interface{}
			if err := json.Unmarshal(resBody, &response); err != nil {
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

		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			path := fmt.Sprintf("%v/", ts.URL)
			res, err := http.Get(path)
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
