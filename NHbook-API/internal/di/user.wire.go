//go:build wireinject

package di

import (
	"database/sql"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/handlers"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitAuthHandler(db *sql.DB, client *redis.Client) (*handlers.AuthHandler, error) {
	wire.Build(
		repositories.NewUserRepository,
		services.NewAuthService,
		handlers.NewAuthHandler,
		services.NewCacheService,
	)

	return nil, nil
}
