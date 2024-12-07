package models

type HealthCheckResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreatePaymentResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CreateAccountResponse struct {
	ID      string `json:"user_id"`
	Message string `json:"message"`
}

type GetPaymentsResponse struct {
	Payments []Payment `json:"payments"`
}
