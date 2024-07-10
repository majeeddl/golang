package main

import (
	"context"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Easier to get running with CORS. Thanks for help @Vindexus and @erkie
var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func main() {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Initialize Socket.IO server
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Subscribe to Redis channel
	pubsub := redisClient.Subscribe(ctx, "socketio_messages")
	defer pubsub.Close()

	// Handle messages from Redis channel in a goroutine
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Fatal(err)
			}
			// Broadcast message to all connected Socket.IO clients
			server.BroadcastToNamespace("/", "message", msg.Payload)
		}
	}()

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("Connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		// Publish message to Redis channel when received from a client
		err := redisClient.Publish(ctx, "socketio_messages", msg).Err()
		if err != nil {
			log.Println("Error publishing message to Redis:", err)
		}
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Disconnected:", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
