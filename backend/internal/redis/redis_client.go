package redis

import (
	"log"

	"github.com/redis/go-redis/v9"
)

// NewRedisClient initializes a new Redis client.
func NewRedisClient(url string) *redis.Client {
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatalf("Error parsing Redis URL: %v", err)
	}

	return redis.NewClient(opts)
}
