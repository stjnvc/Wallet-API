package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func init() {
	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Example: Ping the Redis server to test the connection
	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // Output: PONG
}
