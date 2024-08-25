package main

import (
	"log"
	"net/http"

	"github.com/mahdi-eth/history-storage-service/internal/config/db"
	"github.com/mahdi-eth/history-storage-service/internal/handlers"
	"github.com/mahdi-eth/history-storage-service/internal/subscriber"
)

func main() {
	// Initialize Redis client
	redisClient := db.InitializeRedis()

    // Subscribe to redis channel and save messages
	go subscriber.SubscribeAndSave(redisClient)

	// HTTP routes for storing and retrieving chat history
	http.HandleFunc("/history", handlers.GetHistory)

	log.Println("History storage service started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
