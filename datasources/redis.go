package datasources

import (
	"app/config"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache() Cache {
	client := redis.NewClient(&redis.Options{
		Addr: config.ServiceConfig.Redis.Address,
	})

	return &redisCache{
		client: client,
	}
}

func (c *redisCache) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *redisCache) KeyExists(key string) (bool, error) {
	keys, err := c.client.Keys(context.Background(), key).Result()
	if err != nil {
		return false, err
	}

	if int64(len(keys)) > 0 {
		return true, nil
	}

	return false, nil
}
