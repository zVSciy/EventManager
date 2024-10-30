package models

import "time"

type Payment struct {
	ID               string    `json:"id" bson:"_id,omitempty"`
	UserID           string    `json:"userId" binding:"required"`
	RecipientID      string    `json:"recipientId" binding:"required"`
	Amount           float64   `json:"amount" binding:"required"`
	Currency         string    `json:"currency" binding:"required"`
	PaymentReference string    `json:"paymentReference,omitempty" bson:"paymentReference,omitempty"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"createdAt"`
	ProcessedAt      time.Time `json:"processedAt,omitempty" bson:"processedAt,omitempty"`
}
