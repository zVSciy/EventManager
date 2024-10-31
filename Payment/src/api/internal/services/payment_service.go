package services

import (
	"context"
	"errors"
	"time"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var paymentCollection *mongo.Collection
var idempotencyKeyCollection *mongo.Collection

func InitPaymentService() {
	if db.Client == nil {
		panic("MongoDB client not initialized")
	}
	paymentCollection = db.Client.Database("paymentdb").Collection("payments")
	idempotencyKeyCollection = db.Client.Database("paymentdb").Collection("idempotency_keys")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	util.CreateTTLIndex(ctx, idempotencyKeyCollection)
}

func GetPayment(id string) (models.Payment, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Payment{}, errors.New("invalid_payment_id")
	}

	var payment models.Payment

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = paymentCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&payment)

	if err == nil {
		return payment, nil
	}
	if err == mongo.ErrNoDocuments {
		return models.Payment{}, errors.New("payment_not_found")
	}
	return models.Payment{}, err
}

func CreatePayment(payment models.Payment, idempotencyKeyStr string) (models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	var existingKey models.IdempotencyKey
	err := idempotencyKeyCollection.FindOne(ctx, bson.M{"value": idempotencyKeyStr}).Decode(&existingKey)
	if err == nil {
		return models.Payment{}, errors.New("idempotency_key_error")
	}

	idempotencyKey := models.IdempotencyKey{
		Value:     idempotencyKeyStr,
		CreatedAt: util.Now(),
	}

	_, err = idempotencyKeyCollection.InsertOne(ctx, idempotencyKey)
	if err != nil {
		return models.Payment{}, err
	}

	payment.Status = "initiated"
	payment.CreatedAt = util.Now()

	result, err := paymentCollection.InsertOne(ctx, payment)
	if err != nil {
		return models.Payment{}, err
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		payment.ID = objectID.Hex()
	} else {
		return models.Payment{}, err
	}

	return payment, nil
}
