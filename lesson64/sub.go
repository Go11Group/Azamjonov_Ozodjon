package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})

	// Create a new pubsub client
	pubsub := rdb.Subscribe(ctx, "stockPriceUpdater", "stockPriceMonitor", "stockAnalytics")
	defer pubsub.Close()

	// Wait for confirmation that subscription is created
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Start a goroutine to receive messages
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Printf("Received message from channel %s: %s\n", msg.Channel, msg.Payload)
	}
}
