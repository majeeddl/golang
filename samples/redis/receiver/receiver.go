package main

import (
	"context"
	"fmt"
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

	// Subscribe to Redis channel
	pubsub := redisClient.Subscribe(ctx, "redis_messages")
	defer pubsub.Close()

	// Handle messages from Redis channel in a goroutine
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Fatal(err)
			}
			// Broadcast message to all connected Socket.IO clients
			//

			fmt.Println(msg.Payload)
		}
	}()

	// Wait for all messages from Redis
	for msg := range pubsub.Channel() {
		// Broadcast message to all connected Socket.IO clients
		//
		fmt.Println(msg.Payload)
	}
}
