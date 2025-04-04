package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

// HealthCheck godoc
// @Summary Health Check
// @Tags health
// @Success 200 {object} models.HealthCheckResponse "Service is healthy"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.JSONResponse(w, http.StatusOK, models.HealthCheckResponse{
		Message: "healthy",
	})
}

// GetPayment godoc
// @Summary Get Payment
// @Tags payments
// @Param user_id path string true "User ID"
// @Param id path string true "Payment ID"
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} models.Payment "Payment details"
// @Failure 400 {object} models.ErrorResponse "Invalid payment ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Payment not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/{user_id}/payments/{id} [get]
func GetPayment(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	authHeader := r.Header.Get("Authorization")

	err := util.ValidateAuthHeader(userID, authHeader)
	if err != nil {
		util.JSONResponse(w, http.StatusUnauthorized, models.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	id := r.PathValue("id")

	payment, err := services.GetPayment(id)
	if err == nil {
		if payment.UserID != userID {
			util.JSONResponse(w, http.StatusUnauthorized, models.ErrorResponse{
				Error: "Unauthorized",
			})
			return
		}

		payment.CreatedAt = util.ApplyLocalTZ(payment.CreatedAt)
		if payment.ProcessedAt != nil {
			localProcessedAt := util.ApplyLocalTZ(*payment.ProcessedAt)
			payment.ProcessedAt = &localProcessedAt
		}

		util.JSONResponse(w, http.StatusOK, payment)
		return
	}

	errorResponses := map[string]int{
		"payment_not_found":        http.StatusNotFound,
		"invalid_payment_id":       http.StatusBadRequest,
		"database_not_initialized": http.StatusInternalServerError,
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

// GetPayments godoc
// @Summary Get Payments by Account
// @Tags payments
// @Param user_id path string true "UserID"
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} models.Payment "List of payments"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/{user_id}/payments [get]
func GetPayments(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	authHeader := r.Header.Get("Authorization")

	err := util.ValidateAuthHeader(userID, authHeader)
	if err != nil {
		util.JSONResponse(w, http.StatusUnauthorized, models.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	payments, err := services.GetPayments(userID)
	if err == nil {
		util.JSONResponse(w, http.StatusOK, models.GetPaymentsResponse{
			Payments: payments,
		})
		return
	}

	errorResponses := map[string]int{
		"user_not_found":           http.StatusNotFound,
		"database_not_initialized": http.StatusInternalServerError,
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

// CreatePayment godoc
// @Summary Create Payment
// @Tags payments
// @Param Idempotency-Key header string true "Unique key to prevent duplicate payments"
// @Param payment body models.PaymentRequest true "Payment details"
// @Param Authorization header string true "Authorization token"
// @Success 201 {object} models.CreatePaymentResponse "Payment created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request or missing idempotency key"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 409 {object} models.ErrorResponse "Duplicate idempotency key"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments [post]
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

	authHeader := r.Header.Get("Authorization")

	err := util.ValidateAuthHeader(payment.UserID, authHeader)
	if err != nil {
		util.JSONResponse(w, http.StatusUnauthorized, models.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	// ONLY EUR (TEMPORARY)
	if payment.Currency != "EUR" {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "BAD_CURRENCY_USE_EUR",
		})
		return
	}

	if payment.Amount <= 0 {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "INVALID_AMOUNT",
		})
		return
	}

	createdPayment, err := services.CreatePayment(payment, idempotencyKey)
	if err == nil {
		util.JSONResponse(w, http.StatusCreated, models.CreatePaymentResponse{
			ID:     createdPayment.ID,
			Status: createdPayment.Status,
		})
		return
	}

	errorResponses := map[string]int{
		"idempotency_key_error":    http.StatusConflict,
		"recipient_not_found":      http.StatusNotFound,
		"incompatible_currency":    http.StatusBadRequest,
		"insufficient_funds":       http.StatusBadRequest,
		"database_not_initialized": http.StatusInternalServerError,
	}

	errStr := err.Error()
	if statusCode, exists := errorResponses[errStr]; exists {
		util.JSONResponse(w, statusCode, models.ErrorResponse{
			Error: strings.ToUpper(errStr),
		})
		return
	}

	services.SetPaymentCancelled(payment)

	util.JSONResponse(w, http.StatusInternalServerError, models.ErrorResponse{
		Error: "Internal Server Error",
	})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	util.JSONResponse(w, http.StatusNotFound, models.ErrorResponse{
		Error: "Not Found",
	})
}
