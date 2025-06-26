//go:build wireinject

package wire

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/handlers"
	repo "github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitBookHandler(db *gorm.DB) (*handlers.BookHandler, error) {
	wire.Build(
		repo.NewBookRepository,
		repo.NewAuthorRepository,
		repo.NewCategoryRepository,
		services.NewBookService,
		services.NewCategoryService,
		services.NewAuthorService,
		handlers.NewBookHandler,
	)

	return &handlers.BookHandler{}, nil
}
