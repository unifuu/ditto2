package redis

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	Cli *redis.Client
)

func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return cli
}

// Set key-value pairs with expiration
func Set(key string, value string, expiration time.Duration) error {
	return Cli.Set(ctx, key, value, expiration).Err()
}

// Get value by key
func Get(key string) (string, error) {
	return Cli.Get(ctx, key).Result()
}

// Delete value by key
func Del(key string) error {
	return Cli.Del(ctx, key).Err()
}
