package middlewares

import (
	"time"

	"github.com/NewHorizonIT/nhbook-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestLogger(rootLog *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		start := time.Now()
		reqLog := rootLog.With(
			zap.String("request_id", requestID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)

		// Set logger with request context
		ctx := logger.WithContext(c.Request.Context(), reqLog)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

		// Log request completion
		latency := time.Since(start)
		reqLog.Info("Request completed",
			zap.Int("status", c.Writer.Status()),
			zap.String("latency", latency.String()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}
