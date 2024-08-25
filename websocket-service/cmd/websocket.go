package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mahdi-eth/websocket-service/internal/config/db"
	"github.com/mahdi-eth/websocket-service/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Initialize Redis client
	redisClient := db.InitializeRedis()

	// Start the HTTP server
	err := server.StartServer(redisClient)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
