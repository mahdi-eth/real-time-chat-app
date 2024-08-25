package socket

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// handleConnections manages WebSocket connections and Redis pub/sub
func HandleConnections(w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
    ctx := context.Background()

    // Upgrade initial GET request to a WebSocket
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatalf("Failed to upgrade to WebSocket: %v", err)
    }
    defer ws.Close()

        // Fetch chat history from history service
        history, err := fetchChatHistory()
        if err != nil {
            log.Println("Error fetching chat history:", err)
        } else {
            // Send history to the new client
            for _, msg := range history {
                err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
                if err != nil {
                    log.Println("Error sending chat history to WebSocket client:", err)
                    break
                }
            }
        }

    // Subscribe to Redis channel
    pubsub := redisClient.Subscribe(ctx, "chat_channel")
    defer pubsub.Close()

    // Listen for incoming messages from WebSocket and publish them to Redis
    go handleIncomingMessages(ws, redisClient, ctx)

    // Listen for messages from Redis and send them to WebSocket
    handleRedisMessages(ws, pubsub)
}

func handleIncomingMessages(ws *websocket.Conn, redisClient *redis.Client, ctx context.Context) {
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}
		err = redisClient.Publish(ctx, "chat_channel", string(msg)).Err()
		if err != nil {
			log.Println("Error publishing to Redis:", err)
			break
		}
	}
}

func handleRedisMessages(ws *websocket.Conn, pubsub *redis.PubSub) {
	for msg := range pubsub.Channel() {
		err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		if err != nil {
			log.Println("Error writing WebSocket message:", err)
			break
		}
	}
}

func fetchChatHistory() ([]string, error) {
    resp, err := http.Get(os.Getenv("CHAT_HISTORY"))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var history []string
    err = json.Unmarshal(body, &history)
    if err != nil {
        return nil, err
    }

    return history, nil
}