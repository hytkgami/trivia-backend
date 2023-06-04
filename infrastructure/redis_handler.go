package infrastructure

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hytkgami/trivia-backend/interfaces/repository"
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

func (h *RedisHandler) Ping(ctx context.Context) error {
	return h.client.Ping(ctx).Err()
}

func (h *RedisHandler) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return h.client.Set(ctx, key, value, expiration).Err()
}

func (h *RedisHandler) Get(ctx context.Context, key string) (string, error) {
	return h.client.Get(ctx, key).Result()
}

func (h *RedisHandler) Subscribe(ctx context.Context, channel string) repository.PubSub {
	pubsub := h.client.Subscribe(ctx, channel)
	return &PubSub{
		pubsub: pubsub,
	}
}

func (h *RedisHandler) Publish(ctx context.Context, channel string, message any) error {
	return h.client.Publish(ctx, channel, message).Err()
}

type PubSub struct {
	pubsub *redis.PubSub
}

func (p *PubSub) Close() error {
	return p.pubsub.Close()
}

func (p *PubSub) Channel() <-chan *redis.Message {
	return p.pubsub.Channel()
}
