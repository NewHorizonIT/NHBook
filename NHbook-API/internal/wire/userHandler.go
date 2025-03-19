//go:build wireinject

package wire

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/handlers"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitAuthHandler(db *gorm.DB) (*handlers.AuthHandler, error) {
	wire.Build(
		repositories.NewUserRepository,
		repositories.NewTokenRepository,
		services.NewAuthService,
		handlers.NewAuthHandler,
	)
	return &handlers.AuthHandler{}, nil
}
