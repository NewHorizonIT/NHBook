package initializer

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/middlewares"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Defined router
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Chỉnh domain phù hợp
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Api-Key"},
		AllowCredentials: true,
	}))
	if global.Config.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// Middleware
	r.Use(middlewares.LoggerMiddlerWare())

	r.Use(middlewares.CheckApiKey())
	// Setup Group Router
	newGroupRouter := router.NewRouterGroup
	userRouter := newGroupRouter.UserRouter
	authRouter := newGroupRouter.AuthRouter
	bookRouter := newGroupRouter.BookRouter
	api := r.Group("/api/v1")
	{
		userRouter.SetupRouter(api)
		authRouter.SetupRouter(api)
		bookRouter.SetUpBookRouter(api)
	}

	return r
}
