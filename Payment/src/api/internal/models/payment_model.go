package models

import "time"

type Payment struct {
	UserID           string     `json:"user_id" bson:"user_id" binding:"required"`
	RecipientID      string     `json:"recipient_id" bson:"recipient_id" binding:"required"`
	Amount           float64    `json:"amount" binding:"required"`
	Currency         string     `json:"currency" binding:"required"`
	PaymentReference string     `json:"payment_reference,omitempty" bson:"payment_reference,omitempty"`
	ID               string     `json:"id" bson:"_id,omitempty"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at" bson:"created_at"`
	ProcessedAt      *time.Time `json:"processed_at,omitempty" bson:"processed_at,omitempty"`
}

type PaymentRequest struct {
	UserID           string  `json:"user_id" bson:"user_id" binding:"required"`
	RecipientID      string  `json:"recipient_id" bson:"recipient_id" binding:"required"`
	Amount           float64 `json:"amount" binding:"required"`
	Currency         string  `json:"currency" binding:"required"`
	PaymentReference string  `json:"payment_reference,omitempty" bson:"payment_reference,omitempty"`
}
