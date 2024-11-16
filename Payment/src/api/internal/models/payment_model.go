package models

import "time"

type PaymentBase struct {
	UserID           string  `json:"userId" binding:"required"`
	RecipientID      string  `json:"recipientId" binding:"required"`
	Amount           float64 `json:"amount" binding:"required"`
	Currency         string  `json:"currency" binding:"required"`
	PaymentReference string  `json:"paymentReference,omitempty" bson:"payment_reference,omitempty"`
}

type Payment struct {
	PaymentBase
	ID          string     `json:"id" bson:"_id,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt" bson:"created_at"`
	ProcessedAt *time.Time `json:"processedAt,omitempty" bson:"processed_at,omitempty"`
}

type PaymentRequest struct {
	PaymentBase
}
