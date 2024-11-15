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

	util.JSONResponse(w, http.StatusInternalServerError, models.ErrorResponse{
		Error: "Internal Server Error",
	})
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	payments, err := services.GetPayments(username)
	if err == nil {
		util.JSONResponse(w, http.StatusOK, payments)
	}
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	idempotencyKey := r.Header.Get("Idempotency-Key")
	if idempotencyKey == "" {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "MISSING_IDEMPOTENCY_KEY",
		})
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

	util.JSONResponse(w, http.StatusInternalServerError, models.ErrorResponse{
		Error: "Internal Server Error",
	})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	util.JSONResponse(w, http.StatusNotFound, models.ErrorResponse{
		Error: "Not Found",
	})
}
