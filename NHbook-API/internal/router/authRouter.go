package router

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/di"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar *AuthRouter) SetupRouter(r *gin.RouterGroup) {
	authHandler, err := di.InitAuthHandler(global.MySQL, global.Redis)

	if err != nil {
		global.Logger.Error(err.Error())
		panic("Init AuthHandler Error")
	}

	authRouter := r.Group("/auth")
	{
		// Register
		authRouter.POST("/register", authHandler.Register)
		// Login
		authRouter.POST("/login", authHandler.Login)
		// Logout
		authRouter.POST("/logout", authHandler.Logout)
		// Handle RefreshToken
		authRouter.POST("/refresh-token", authHandler.HandleRefreshToken)
		// Get info user
		authRouter.GET("me", authHandler.GetInfoUser)
	}
}
