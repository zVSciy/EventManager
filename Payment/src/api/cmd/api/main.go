package main

import (
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/zVSciy/EventManager/Payment/docs"

	db "github.com/zVSciy/EventManager/Payment/internal/database"
	"github.com/zVSciy/EventManager/Payment/internal/handlers"
	"github.com/zVSciy/EventManager/Payment/internal/middleware"
	"github.com/zVSciy/EventManager/Payment/internal/services"
	"github.com/zVSciy/EventManager/Payment/internal/util"
)

// @title Payment Service API
// @version 1.0
// @description API for managing payments
// @host reiner.gg
// @BasePath /api/v1
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

	mux.HandleFunc("GET /docs", handlers.GetDocs)
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/docs"),
	))

	api := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", api))

	api.HandleFunc("GET /health", handlers.HealthCheck)
	api.HandleFunc("GET /accounts/{user_id}/payments", handlers.GetPayments)
	api.HandleFunc("GET /payments/{id}", handlers.GetPayment)
	api.HandleFunc("POST /payments", handlers.CreatePayment)

	mux.HandleFunc("/", handlers.NotFound)

	middlewareChain := middleware.CreateChain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    PORT,
		Handler: middlewareChain(mux),
	}

	log.Printf("Starting server on %s", PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
