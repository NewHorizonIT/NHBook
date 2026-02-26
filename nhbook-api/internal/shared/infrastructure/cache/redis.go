package cache

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/redis/go-redis/v9"
)

// Define Errors
var (
	ErrKeyNotFound = errors.New("key not found")
	ErrNilValue    = errors.New("nil value")
)

// Implement cache with Redis
type RedisCache struct {
	client *redis.Client
	prefix string
}

// Initialize Redis cache
func NewRedisCache(cfg config.RedisConfig) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
	})

	// Test connection
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client: client,
		prefix: cfg.Prefix,
	}, nil
}

// Helper to add prefix to keys
func (r *RedisCache) prefixedKey(key string) string {
	return r.prefix + key
}

// =========== Implement ICache Methods ===========

// Oparations
func (r *RedisCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, r.prefixedKey(key), data, ttl).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) (any, error) {
	data, err := r.client.Get(ctx, r.prefixedKey(key)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrKeyNotFound
		}
		return nil, err
	}
	var dest any
	if err := json.Unmarshal([]byte(data), &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, r.prefixedKey(key)).Err()
}

func (r *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, r.prefixedKey(key)).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Check health
func (r *RedisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

// Close the cache connection
func (r *RedisCache) Close(ctx context.Context) error {
	return r.client.Close()
}
