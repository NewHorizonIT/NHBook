// @title           NHBook API
// @version         1.0
// @description     API cho hệ thống bán sách
// @termsOfService  https://example.com/terms/

// @contact.name   Anh Quan
// @contact.email  nguyenanhquandev@gmail.com

// @host      localhost:5555
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"context"

	_ "github.com/NewHorizonIT/nhbook-api/docs"
	authModule "github.com/NewHorizonIT/nhbook-api/internal/modules/auth"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/infrastructure/cache"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/infrastructure/database"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/middlewares"
	"github.com/NewHorizonIT/nhbook-api/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	// Read configfuration
	cnf, err := config.Load(&config.LoadConfig{
		Path: "./configs",
		Name: "config.example",
		Type: "yaml",
	})
	if err != nil {
		panic(err)
	}

	// Initialize logger
	log, err := logger.InitLogger(&cnf.Log)
	if err != nil {
		panic(err)
	}
	log.App.Info("Logger initialized successfully")
	defer logger.Sync(log)

	// Init database
	db, err := database.InitDB(&cnf.Db)
	if err != nil {
		log.Error.Error("Failed to initialize database", zap.Error(err))
	}

	// Ping database
	if err := db.Ping(); err != nil {
		log.Error.Error("Failed to connect to database", zap.Error(err))
	}

	defer db.Close()
	log.App.Info("Database connection established successfully")

	log.App.Info("Starting server...",
		zap.String("host", cnf.App.Host),
		zap.Int("port", cnf.App.Port),
		zap.String("app_name", cnf.App.Name),
		zap.Bool("env", cnf.App.DevelopmentMode),
	)

	// Initialize cache with Redis
	redis, err := cache.NewRedisCache(cnf.Redis)
	if err != nil {
		log.Error.Error("Failed to initialize Redis cache", zap.Error(err))
	}
	defer redis.Close(context.Background())

	// Register routes
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	// Check API health
	apiGroup.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User requests middleware
	apiGroup.Use(middlewares.RequestLogger(log.Request))

	// ========== Initialize Modules ==========
	// Auth Module
	auth := authModule.NewAuthModule(db, *cnf)
	auth.RegisterRoutes(apiGroup)
	log.App.Info("Auth module initialized successfully")

	// Start server
	router.Run(":5555")
}
