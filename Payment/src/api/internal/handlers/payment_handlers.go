package handlers

import (
	"encoding/json"
	"net/http"

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
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	payment.CreatedAt = util.ApplyLocalTZ(payment.CreatedAt)
	if payment.ProcessedAt != nil {
		localProcessedAt := util.ApplyLocalTZ(*payment.ProcessedAt)
		payment.ProcessedAt = &localProcessedAt
	}

	util.JSONResponse(w, http.StatusOK, payment)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	idempotencyKey := r.Header.Get("Idempotency-Key")

	if idempotencyKey == "" {
		http.Error(w, "Idempotency-Key is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	createdPayment, err := services.CreatePayment(payment, idempotencyKey)
	if err == nil {
		util.JSONResponse(w, http.StatusOK, createdPayment)
		return
	}

	if err.Error() == "idempotency_key_error" {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "404 Not Found", http.StatusNotFound)
}
