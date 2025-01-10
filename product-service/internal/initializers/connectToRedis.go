package initializers

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() (*redis.Client, error) {
	var client *redis.Client
	var err error

	ctx := context.Background()
	url := os.Getenv("REDIS_URL")

	for i := 0; i < 5; i++ {
		client = redis.NewClient(&redis.Options{
			Addr: url,
		})

		_, err = client.Ping(ctx).Result()
		if err == nil {
			log.Println("Connected to Redis")
			return client, nil
		}
		log.Printf("Failed to connect to Redis (attempt %d): %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, errors.New("failed to connect to Redis after 5 attempts")
}
