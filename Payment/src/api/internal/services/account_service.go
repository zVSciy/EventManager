package services

import (
	"context"
	"errors"
	"time"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/models"
	"github.com/zVSciy/EventManager/Payment/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection

func InitAccountService() {
	if db.Client == nil {
		panic("MongoDB client not initialized")
	}
	accountCollection = db.Client.Database("paymentdb").Collection("accounts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	util.CreateUniqueAccountIndex(ctx, accountCollection)
}

func CreateAccount(account models.Account) (models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := accountCollection.FindOne(ctx, bson.M{"user_id": account.ID}).Err(); err == nil {
		return models.Account{}, errors.New("account_exists")
	}

	account.Balance = 100
	account.Currency = "EUR"
	account.CreatedAt = util.Now()

	_, err := accountCollection.InsertOne(ctx, account)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}
