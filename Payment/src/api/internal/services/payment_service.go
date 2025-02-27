package services

import (
	"context"
	"errors"
	"log"
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
	if err := util.CheckCollectionInit(paymentCollection); err != nil {
		return models.Payment{}, errors.New("database_not_initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var payment models.Payment
	err = paymentCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&payment)
	if err == nil {
		return payment, nil
	}

	if err == mongo.ErrNoDocuments {
		return models.Payment{}, errors.New("payment_not_found")
	}

	return models.Payment{}, err
}

func GetPayments(userId string) ([]models.Payment, error) {
	if err := util.CheckCollectionInit(paymentCollection); err != nil {
		return []models.Payment{}, errors.New("database_not_initialized")
	}

	var account models.Account
	payments := []models.Payment{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := accountCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.Payment{}, errors.New("user_not_found")
		}
		return []models.Payment{}, err
	}

	var testPayment models.Payment
	paymentCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&testPayment)
	log.Println(testPayment)

	cur, err := paymentCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return []models.Payment{}, err
	}

	if err = cur.All(ctx, &payments); err != nil {
		return []models.Payment{}, err
	}

	return payments, nil
}

func CreatePayment(payment models.Payment, idempotencyKeyStr string) (models.Payment, error) {
	if err := util.CheckCollectionInit(paymentCollection); err != nil {
		return models.Payment{}, errors.New("database_not_initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingKey models.IdempotencyKey
	err := idempotencyKeyCollection.FindOne(ctx, bson.M{"value": idempotencyKeyStr}).Decode(&existingKey)
	if err == nil {
		return models.Payment{}, errors.New("idempotency_key_error")
	}

	count, err := accountCollection.CountDocuments(ctx, bson.M{"user_id": payment.RecipientID})
	if err != nil {
		return models.Payment{}, err
	}
	if count == 0 {
		return models.Payment{}, errors.New("recipient_not_found")
	}

	var account models.Account
	err = accountCollection.FindOne(ctx, bson.M{"user_id": payment.UserID}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// ACCOUNT CREATION TEMPORARY SOLUTION
			account = models.Account{
				ID:        payment.UserID,
				Balance:   100,
				Currency:  "EUR",
				CreatedAt: util.Now(),
			}

			_, err = accountCollection.InsertOne(ctx, account)
			if err != nil {
				return models.Payment{}, err
			}
		} else {
			return models.Payment{}, err
		}
	}

	idempotencyKey := models.IdempotencyKey{
		Value:     idempotencyKeyStr,
		CreatedAt: util.Now(),
	}

	_, err = idempotencyKeyCollection.InsertOne(ctx, idempotencyKey)
	if err != nil {
		return models.Payment{}, err
	}

	if payment.Currency != account.Currency {
		return models.Payment{}, errors.New("incompatible_currency")
	}

	if payment.Amount > account.Balance {
		return models.Payment{}, errors.New("insufficient_funds")
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

	insertedID := result.InsertedID.(primitive.ObjectID)
	payment.ID = insertedID.Hex()

	if _, err := accountCollection.UpdateOne(
		ctx,
		bson.M{"user_id": payment.UserID},
		bson.M{"$inc": bson.M{"balance": -payment.Amount}},
	); err != nil {
		return models.Payment{}, err
	}

	if _, err := accountCollection.UpdateOne(
		ctx,
		bson.M{"user_id": payment.RecipientID},
		bson.M{"$inc": bson.M{"balance": payment.Amount}},
	); err != nil {
		return models.Payment{}, err
	}

	if _, err := paymentCollection.UpdateOne(
		ctx,
		bson.M{"_id": insertedID},
		bson.M{"$set": bson.M{"status": "processed"}},
	); err != nil {
		return models.Payment{}, err
	}
	payment.Status = "processed"

	return payment, nil
}

func SetPaymentCancelled(payment models.Payment) error {
	paymentID, err := primitive.ObjectIDFromHex(payment.ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := paymentCollection.UpdateOne(
		ctx,
		bson.M{"_id": paymentID},
		bson.M{"$set": bson.M{"status": "cancelled"}},
	); err != nil {
		return err
	}

	return nil
}
