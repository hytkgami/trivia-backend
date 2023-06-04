package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisHandler interface {
	Close() error
	Ping(ctx context.Context) error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Subscribe(ctx context.Context, channel string) PubSub
	Publish(ctx context.Context, channel string, message any) error
}

type PubSub interface {
	Close() error
	Channel() <-chan *redis.Message // TODO: redis依存を剥がしたい
}
