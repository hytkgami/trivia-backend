package repository

import (
	"context"
	"time"
)

type RedisHandler interface {
	Close() error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}