package main

import (
	"log"
	"os"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/handlers"
	"github.com/zVSciy/EventManager/Payment/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://db-payment:27017"
	}

	log.Println("Initializing MongoDB client...")
	db.Init(mongoURI)
	log.Println("MongoDB client initialized successfully")

	log.Println("Initializing Payment service")
	services.InitPaymentService()
	log.Println("Payment service initialized successfully")

	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)
	r.POST("/payments", handlers.CreatePayment)
	// r.POST("/payments/:id/process", handlers.ProcessPayment)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
