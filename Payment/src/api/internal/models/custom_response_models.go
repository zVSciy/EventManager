package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreatePaymentResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
