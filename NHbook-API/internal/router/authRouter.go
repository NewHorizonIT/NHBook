package router

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/wire"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar *AuthRouter) SetupRouter(r *gin.RouterGroup) {
	authHandler, err := wire.InitAuthHandler(global.MySQL)
	if err != nil {
		global.Logger.Error(err.Error())
		panic("Init Authhandler Error")
	}
	authRouter := r.Group("/auth")
	{
		// Register
		authRouter.POST("/register", authHandler.Register)
		// Login
		authRouter.POST("/login", authHandler.Login)
		// Logout
		authRouter.POST("/logout", authHandler.Logout)
		// Handle Refreshtoken
		authRouter.POST("/refresh-token", authHandler.HandleRefreshToken)
	}
}
