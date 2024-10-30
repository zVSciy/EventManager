package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, models.HealthCheckResponse{
		Message: "healthy",
	})
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdPayment := services.CreatePayment(payment)
	util.JSON(w, http.StatusOK, createdPayment)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 Not Found")
}
