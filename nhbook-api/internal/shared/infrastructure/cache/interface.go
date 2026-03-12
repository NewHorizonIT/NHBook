package cache

import (
	"context"
	"time"
)

// ICache defines the interface for cache operations
type ICache interface {
	// Check health
	Ping(ctx context.Context) error

	// Operations
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Get(ctx context.Context, key string) (any, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)

	// Close the cache connection
	Close(ctx context.Context) error
}
