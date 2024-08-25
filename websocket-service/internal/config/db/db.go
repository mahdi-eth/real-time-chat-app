package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

// InitializeRedis initializes the Redis client
func InitializeRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	
	return RedisClient
}