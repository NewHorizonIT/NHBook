package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddlerware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("%s | %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
