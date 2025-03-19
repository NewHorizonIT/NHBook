package initializer

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Defind router
	r := gin.Default()
	r.Use(gin.Recovery())
	if global.Config.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// Setup Group Router
	newGroupRouter := router.NewRoutergroup
	userRouter := newGroupRouter.UserRouter
	authRouter := newGroupRouter.AuthRouter
	api := r.Group("/api/v1")
	{

		userRouter.SetupRouter(api)
		authRouter.SetupRouter(api)
	}

	return r
}
