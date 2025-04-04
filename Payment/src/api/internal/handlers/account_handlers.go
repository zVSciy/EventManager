package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

// GetAccount godoc
// @Summary Get Account Details
// @Tags accounts
// @Param user_id path string true "User ID"
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} models.Account "Account details"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Account not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/{user_id} [get]
func GetAccount(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	authHeader := r.Header.Get("Authorization")

	err := util.ValidateAuthHeader(userID, authHeader)
	if err != nil {
		util.JSONResponse(w, http.StatusUnauthorized, models.ErrorResponse{
			Error: "Unauthorized",
		})
		return
	}

	account, err := services.GetAccount(userID)
	if err == nil {
		account.CreatedAt = util.ApplyLocalTZ(account.CreatedAt)
		util.JSONResponse(w, http.StatusOK, account)
		return
	}

	errorResponses := map[string]int{
		"account_not_found":        http.StatusNotFound,
		"database_not_initialized": http.StatusInternalServerError,
	}

	errStr := err.Error()
	if statusCode, exists := errorResponses[errStr]; exists {
		util.JSONResponse(w, statusCode, models.ErrorResponse{
			Error: strings.ToUpper(errStr),
		})
	}
}

// CreateAccount godoc
// @Summary Create Account
// @Tags accounts
// @Param account body models.AccountRequest true "User ID"
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} models.CreateAccountResponse "Account created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		util.JSONResponse(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "Bad Request",
		})
		return
	}

	authHeader := r.Header.Get("Authorization")

	err := util.ValidateAuthHeader(account.ID, authHeader)
	if err != nil {
		util.JSONResponse(w, http.StatusUnauthorized, models.ErrorResponse{
			Error: "Unauthorized",
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
		"account_exists":           http.StatusConflict,
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
