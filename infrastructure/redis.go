package infrastructure

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	client *redis.Client
}

func NewRedisHandler() *RedisHandler {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	db, err := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 1)
	if err != nil {
		db = 0
	}
	return &RedisHandler{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       int(db),
		}),
	}
}

func (h *RedisHandler) Close() error {
	return h.client.Close()
}

func (h *RedisHandler) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return h.client.Set(ctx, key, value, expiration).Err()
}

func (h *RedisHandler) Get(ctx context.Context, key string) (string, error) {
	return h.client.Get(ctx, key).Result()
}
