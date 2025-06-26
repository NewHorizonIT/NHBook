//go:build wireinject

package wire

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/handlers"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitCategoryHanler(db *gorm.DB) (*handlers.CategoryHandler, error) {
	wire.Build(
		repositories.NewCategoryRepository,
		services.NewCategoryService,
		handlers.NewCategoryHandler,
	)

	return &handlers.CategoryHandler{}, nil
}
