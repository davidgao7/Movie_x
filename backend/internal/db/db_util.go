package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"          // Import the pq driver
	"github.com/redis/go-redis/v9" // Assuming you are using the go-redis package
)

// getRedisClient initializes and returns a Redis client using a URL.
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

// getPostgresClient initializes and returns a PostgreSQL client.
func GetPostgresClient(postgresURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
