package middlewares

import (
	"fmt"
	"net/http"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Get api key
		apiKeyOfClient := c.GetHeader(global.HEADER_API_KEY)
		fmt.Printf("APIKEY :: %v\n", apiKeyOfClient)

		if apiKeyOfClient == "" {
			utils.WriteError(c, http.StatusUnauthorized, "Missing API KEY")
			c.Abort()
			return
		}

		// Step 2: Compare ApiKey
		apiKeyOfServer := global.Config.ApiKey

		if apiKeyOfServer != apiKeyOfClient {
			utils.WriteError(c, http.StatusUnauthorized, "Api key Invalid")
			c.Abort()
			return
		}

		c.Next()
	}
}
