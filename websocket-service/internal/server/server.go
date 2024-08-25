package server

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	socket "github.com/mahdi-eth/websocket-service/internal/websocket"
)

// StartServer initializes and starts the HTTP server
func StartServer(redisClient *redis.Client) error {
    // HTTP route for WebSocket connection
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
       socket.HandleConnections(w, r, redisClient)
    })

    log.Println("WebSocket service started on :8080")
    return http.ListenAndServe(":8080", nil)
}
