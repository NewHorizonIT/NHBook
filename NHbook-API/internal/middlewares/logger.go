package middlewares

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddlerWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request"})
			return
		}

		log.Printf("[GIN REQUEST] %s %s\nBody: %s\n", c.Request.Method, c.Request.URL.Path, string(bodyBytes))

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}
