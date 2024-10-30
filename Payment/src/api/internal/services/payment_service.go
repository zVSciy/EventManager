package services

import (
	"context"
	"log"
	"time"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/models"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var paymentCollection *mongo.Collection

func InitPaymentService() {
	if db.Client == nil {
		panic("MongoDB client not initialized")
	}
	paymentCollection = db.Client.Database("paymentdb").Collection("payments")
}

func CreatePayment(payment models.Payment) models.Payment {
	loc, err := time.LoadLocation("Europe/Vienna") // TODO: Get TZ from env
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}
	payment.Status = "initiated"
	payment.CreatedAt = time.Now().In(loc)

	log.Printf("Payment: %+v\n", payment)

	result, err := paymentCollection.InsertOne(context.Background(), payment)
	if err != nil {
		panic(err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		payment.ID = oid.Hex()
	} else {
		log.Fatalf("Failed")
	}

	return payment
}
