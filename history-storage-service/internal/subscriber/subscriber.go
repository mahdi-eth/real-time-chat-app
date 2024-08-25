package subscriber

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/mahdi-eth/history-storage-service/internal/config/db"
)

func SubscribeAndSave(redisClient *redis.Client) {
    pubsub := redisClient.Subscribe(db.Ctx, "chat_channel")
    defer pubsub.Close()

    for msg := range pubsub.Channel() {
        err := redisClient.LPush(db.Ctx, "chat_history", msg.Payload).Err()
        if err != nil {
            log.Println("Failed to save message:", err)
        } else {
            log.Println("Message saved:", msg.Payload)
        }
    }
}