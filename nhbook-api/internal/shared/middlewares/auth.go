package middlewares

import (
	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/NewHorizonIT/nhbook-api/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cnf config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing"})
			return
		}
		// Slice the "Bearer " prefix
		if len(tokenString) < 7 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format"})
			return
		}
		tokenString = tokenString[len("Bearer "):]
		// Get secret key from config
		secretKey := cnf.Secret
		// Validate the token
		claims, err := auth.ParseJWT(tokenString, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}
		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}
