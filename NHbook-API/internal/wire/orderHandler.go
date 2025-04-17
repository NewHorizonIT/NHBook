//go:build wireinject

package wire

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/handlers"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitOrderHandler(db *gorm.DB, rd *redis.Client) (*handlers.OrderHandler, error) {
	wire.Build(
		repositories.NewOrderRepository,
		repositories.NewUserRepository,
		repositories.NewBookRepository,
		repositories.NewCartRepository,
		services.NewOrderService,
		handlers.NewOrderHandler,
	)

	return &handlers.OrderHandler{}, nil
}
