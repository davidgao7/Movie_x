package db

import (
	"context"

	"github.com/redis/go-redis/v9" // Assuming you are using the go-redis package
)

// GetRedisClient initializes and returns a Redis client using a URL.
func GetRedisClient(redisURL string) (*redis.Client, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(options)
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
