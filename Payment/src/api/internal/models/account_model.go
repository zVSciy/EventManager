package models

import "time"

type Account struct {
	ID        string    `json:"user_id" bson:"user_id" binding:"required"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency" bson:"currency"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}

type AccountRequest struct {
	ID string `json:"user_id" bson:"user_id"`
}
