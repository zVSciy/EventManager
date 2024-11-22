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
var accountCollection *mongo.Collection
var idempotencyKeyCollection *mongo.Collection

func InitPaymentService() {
	if db.Client == nil {
		panic("MongoDB client not initialized")
	}
	paymentCollection = db.Client.Database("paymentdb").Collection("payments")
	accountCollection = db.Client.Database("paymentdb").Collection("accounts")
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

func GetPayments(userId string) ([]models.Payment, error) {
	var account models.Account
	var payments []models.Payment

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := accountCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []models.Payment{}, errors.New("user_not_found")
		}
		return []models.Payment{}, err
	}

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
	if err == nil {
		// ACCOUNT CREATION TEMPORARY SOLUTION
		account := models.Account{
			ID:        payment.UserID,
			Balance:   100,
			Currency:  "EUR",
			CreatedAt: util.Now(),
		}
		_, err = accountCollection.InsertOne(ctx, account)
		if err != nil {
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

	session, err := db.Client.StartSession()
	if err != nil {
		return models.Payment{}, err
	}
	if err = session.StartTransaction(); err != nil {
		return models.Payment{}, err
	}

	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if _, err := accountCollection.UpdateOne(
			sc,
			bson.M{"user_id": payment.UserID},
			bson.M{"$inc": bson.M{"balance": -payment.Amount}},
		); err != nil {
			return err
		}
		if _, err := accountCollection.UpdateOne(
			sc,
			bson.M{"user_id": payment.RecipientID},
			bson.M{"$inc": bson.M{"balance": payment.Amount}},
		); err != nil {
			return err
		}

		if _, err := paymentCollection.UpdateOne(
			sc,
			bson.M{"_id": payment.ID},
			bson.M{"$set": bson.M{"status": "processed"}},
		); err != nil {
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}
		return nil
	}); err != nil {
		_ = session.AbortTransaction(ctx)
		return models.Payment{}, err
	}
	session.EndSession(ctx)

	return payment, nil
}

func SetPaymentCancelled(payment models.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := paymentCollection.UpdateOne(
		ctx,
		bson.M{"_id": payment.ID},
		bson.M{"$set": bson.M{"status": "cancelled"}},
	); err != nil {
		return err
	}

	return nil
}
