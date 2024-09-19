package db

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetClient(url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		fmt.Println("Error parsing redis url: ", err)
		return nil, err
	}

	return redis.NewClient(opts), nil
}
