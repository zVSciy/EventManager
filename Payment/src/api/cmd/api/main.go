package main

import (
	"os"
	"api/internal/database"
	"api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://db-payment:27017"
	}

	db.ConnectMongoDB(mongoURI)

	r := gin.Default()

	r.POST("/payments", handlers.CreatePayment)
	r.POST("/payments/:id/process", handlers.ProcessPayment)

	r.Run(":3000")
}