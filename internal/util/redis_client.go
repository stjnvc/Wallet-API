package util

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient() error {
	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password
		DB:       0,                // Redis database index
	})

	// Ping the Redis server to check the connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}
