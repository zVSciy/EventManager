package models

import "time"

type IdempotencyKey struct {
	Value     string    `bson:"value"`
	CreatedAt time.Time `bson:"created_at"`
}
