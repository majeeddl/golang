package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := redisClient.Publish(ctx, "redis_messages", "hello").Err()
	if err != nil {
		log.Println("Error publishing message to Redis:", err)
	}
}
