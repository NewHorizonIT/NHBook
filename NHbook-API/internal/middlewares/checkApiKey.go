package middlewares

import (
	"net/http"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

func checkApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(global.HEADER_API_KEY)

		if apiKey == "" {
			utils.WriteError(c, http.StatusForbidden, "Missing API KEY")
		}

		c.Next()
	}
}
