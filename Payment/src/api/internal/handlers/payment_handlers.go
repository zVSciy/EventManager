package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.JSONResponse(w, http.StatusOK, models.HealthCheckResponse{
		Message: "healthy",
	})
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	payment, err := services.GetPayment(id)
	if err == nil {
		payment.CreatedAt = util.ApplyLocalTZ(payment.CreatedAt)
		if payment.ProcessedAt != nil {
			localProcessedAt := util.ApplyLocalTZ(*payment.ProcessedAt)
			payment.ProcessedAt = &localProcessedAt
		}

		util.JSONResponse(w, http.StatusOK, payment)
		return
	}

	errorResponses := map[string]int{
		"payment_not_found":  http.StatusNotFound,
		"invalid_payment_id": http.StatusBadRequest,
	}

	errStr := err.Error()
	if statusCode, exists := errorResponses[errStr]; exists {
		util.JSONResponse(w, statusCode, models.ErrorResponse{
			Error: strings.ToUpper(errStr),
		})
		return
	}

	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	idempotencyKey := r.Header.Get("Idempotency-Key")
	if idempotencyKey == "" {
		http.Error(w, "Idempotency-Key is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "Bad Request",
		})
		return
	}

	createdPayment, err := services.CreatePayment(payment, idempotencyKey)
	if err == nil {
		util.JSONResponse(w, http.StatusOK, models.CreatePaymentResponse{
			ID:     createdPayment.ID,
			Status: createdPayment.Status,
		})
		return
	}

	errStr := err.Error()
	if errStr == "idempotency_key_error" {
		util.JSONResponse(w, http.StatusConflict, models.ErrorResponse{
			Error: strings.ToUpper(errStr),
		})
		return
	}

	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "404 Not Found", http.StatusNotFound)
}
