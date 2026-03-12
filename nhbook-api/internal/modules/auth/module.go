package auth

import (
	"database/sql"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/infrastructure"
	presentation "github.com/NewHorizonIT/nhbook-api/internal/modules/auth/presentation/http"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/gin-gonic/gin"
)

// AuthModule quản lý toàN bộ auth module
type AuthModule struct {
	handler *presentation.AuthHandler
}

// NewAuthModule khởi tạo auth module với tất cả dependencies
func NewAuthModule(db *sql.DB, cfg config.Config) *AuthModule {
	// 1. Infrastructure Layer - Repository
	userRepo := infrastructure.NewUserRepository(db)

	// 2. Application Layer - UseCases
	createUserUsecase := usecase.NewCreateUserUsecase(userRepo, cfg.JWT)
	loginUsecase := usecase.NewLoginUsecase(userRepo, cfg.JWT)

	// 3. Presentation Layer - Handler
	handler := presentation.NewAuthHandler(createUserUsecase, loginUsecase)

	return &AuthModule{
		handler: handler,
	}
}

// RegisterRoutes đăng ký routes cho module
func (m *AuthModule) RegisterRoutes(router *gin.RouterGroup) {
	presentation.RegisterAuthRoutes(router, m.handler)
}

// GetRepository trả về repository instance (nếu cần dùng ở nơi khác)
func (m *AuthModule) GetHandler() *presentation.AuthHandler {
	return m.handler
}
