package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type ICache interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type Cache struct {
	client *redis.Client
}

func NewCacheService(client *redis.Client) ICache {
	return &Cache{
		client: client,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}
