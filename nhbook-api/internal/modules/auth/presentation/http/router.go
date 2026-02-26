package presentation

import "github.com/gin-gonic/gin"

// RegisterAuthRoutes đăng ký các routes cho auth module
func RegisterAuthRoutes(r *gin.RouterGroup, handler *AuthHandler) {
	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signup", handler.SignUp)
		authRouter.POST("/login", handler.Login)
	}
}
