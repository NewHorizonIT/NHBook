package initializer

import (
	"time"

	_ "github.com/NguyenAnhQuan-Dev/NKbook-API/docs"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/middlewares"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// Defined router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://localhost:3030"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "x-api-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

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
	cartRouter := newGroupRouter.CartRouter
	orderRouter := newGroupRouter.OrderRouter
	categoryRouter := newGroupRouter.CategoryRouter
	api := r.Group("/api/v1")
	{
		userRouter.SetupRouter(api)
		authRouter.SetupRouter(api)
		bookRouter.SetUpBookRouter(api)
		cartRouter.SetUpRouter(api)
		orderRouter.SetupRouter(api)
		categoryRouter.SetUpRouter(api)
	}

	return r
}
