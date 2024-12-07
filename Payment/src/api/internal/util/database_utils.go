package util

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTTLIndex(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "created_at", Value: 1},
		},
		Options: options.Index().SetExpireAfterSeconds(2592000), // 30 days
	})
	return err
}

func CreateUniqueAccountIndex(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"user_id": 1},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func CheckCollectionInit(collection *mongo.Collection) error {
	if collection == nil {
		return errors.New("database not initialized")
	}
	return nil
}
