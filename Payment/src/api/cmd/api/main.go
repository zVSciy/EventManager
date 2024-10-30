package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/handlers"
	"github.com/zVSciy/EventManager/Payment/internal/services"
)

func main() {
	MONGO_URI := getenv("MONGO_URI", "mongodb://db-payment:27017")
	PORT := fmt.Sprintf(":%s", getenv("PORT", "3000"))

	log.Println("Initializing MongoDB client...")
	db.Init(MONGO_URI)
	log.Println("MongoDB client initialized successfully")

	log.Println("Initializing Payment service...")
	services.InitPaymentService()
	log.Println("Payment service initialized successfully")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthCheck)
	mux.HandleFunc("POST /payments", handlers.CreatePayment)

	mux.HandleFunc("/", handlers.NotFound)

	server := http.Server{
		Addr:    PORT,
		Handler: mux,
	}

	log.Printf("Starting server on port %s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
