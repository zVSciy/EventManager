package util

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTTLIndex(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "created_at", Value: 1},
		},
		Options: options.Index().SetExpireAfterSeconds(2592000), //expire 30 days
	})
	return err
}
