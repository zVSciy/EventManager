package models

import "time"

type Account struct {
	ID        string    `json:"userId" bson:"user_id" binding:"required"`
	Balance   float64   `json:"balance" binding:"required"`
	Currency  string    `json:"currency" binding:"required"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}
