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

func IniCartHandler(rd *redis.Client, db *gorm.DB) (*handlers.CartHandler, error) {
	wire.Build(
		repositories.NewCartRepository,
		repositories.NewBookRepository,
		repositories.NewCategoryRepository,
		repositories.NewAuthorRepository,
		services.NewCategoryService,
		services.NewAuthorService,
		services.NewCartService,
		services.NewBookService,
		handlers.NewCartHandler,
	)

	return &handlers.CartHandler{}, nil
}
