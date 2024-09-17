package dbutil

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func getClient(url string) *redis.Client {
	opts, err := redis.ParseURL(url)
	if err != nil {
		fmt.Println("Error parsing redis url: ", err)
		return nil
	}

	return redis.NewClient(opts)
}
