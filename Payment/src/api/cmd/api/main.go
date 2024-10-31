package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/handlers"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

func main() {
	MONGO_URI := util.Getenv("MONGO_URI", "mongodb://db-payment:27017")
	PORT := fmt.Sprintf(":%s", util.Getenv("PORT", "3000"))
	TZ := util.Getenv("TZ", "Europe/Vienna")

	log.Println("Initializing Timezone...")
	util.InitTimezone(TZ)

	log.Println("Initializing MongoDB client...")
	db.Init(MONGO_URI)

	log.Println("Initializing Payment service...")
	services.InitPaymentService()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthCheck)
	mux.HandleFunc("GET /payments/{id}", handlers.GetPayment)
	mux.HandleFunc("POST /payments", handlers.CreatePayment)

	mux.HandleFunc("/", handlers.NotFound)

	server := http.Server{
		Addr:    PORT,
		Handler: mux,
	}

	log.Printf("Starting server on %s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
