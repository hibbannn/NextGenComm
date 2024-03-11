package contract

import (
	"context"
	"time"
)

type CacheContract interface {
	GetUserDetailFromRedisSession(ctx context.Context, sessionKey string) (string, error)
}

type RedisContract interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
