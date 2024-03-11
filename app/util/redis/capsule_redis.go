package redis

import (
	"context"
	"fmt"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	config domain.CacheConfig
	client *redis.Client
}

func NewRedis(cfg domain.CacheConfig) *Redis {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@%s:%s/%d",
		"redis",
		cfg.Password, cfg.Host, cfg.Port, cfg.Db,
	))
	if err != nil {
		return nil
	}

	// Menyesuaikan opsi koneksi dengan konfigurasi
	opt.PoolSize = cfg.MaxConn
	opt.MinIdleConns = cfg.MaxIdle
	opt.ConnMaxLifetime = time.Duration(cfg.MaxAge) * time.Second
	opt.ConnMaxIdleTime = time.Duration(cfg.Timeout) * time.Second

	// Membuat klien Redis
	client := redis.NewClient(opt)

	// Melakukan ping untuk memastikan koneksi berhasil
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil
	}

	return &Redis{
		client: client,
		config: cfg,
	}
}
