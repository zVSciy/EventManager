package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

// CreateAccount godoc
// @Summary Create Account
// @Tags accounts
// @Param account body models.AccountRequest true "Account ID"
// @Success 200 {object} models.CreateAccountResponse "Account created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Router /accounts [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "Bad Request",
		})
		return
	}

	createdAccount, err := services.CreateAccount(account)
	if err == nil {
		util.JSONResponse(w, http.StatusCreated, models.CreateAccountResponse{
			ID:      createdAccount.ID,
			Message: "Account created successfully",
		})
		return
	}

	errorResponses := map[string]int{
		"account_exists": http.StatusConflict,
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
