package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

// CreateAccount godoc
// @Summary Create Account
// @Tags accounts
// @Param account body models.AccountRequest true "Account ID"
// @Success 200 {object} models.CreateAccountResponse "Account created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Router /accounts/{user_id} [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.AccountRequest

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "Bad Request",
		})
		return
	}
}
