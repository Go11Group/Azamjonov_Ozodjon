package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectDB() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	err := client.Ping(context.Background()).Err()

	return client, err
}
