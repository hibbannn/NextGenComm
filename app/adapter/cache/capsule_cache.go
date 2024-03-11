package cache

import "github.com/hibbannn/grpc-http-boilerplate/app/core/contract"

type Cache struct {
	redis contract.RedisContract
}

func NewCache(redis contract.RedisContract) *Cache {
	return &Cache{
		redis: redis,
	}
}
