package middlewares

import (
	"fmt"
	"net/http"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(global.HEADER_API_KEY)
		fmt.Print("APIKEY :: ", apiKey)

		if apiKey == "" {
			utils.WriteError(c, http.StatusForbidden, "Missing API KEY")
			c.Abort()
			return
		}
		apiKeyRepo := repositories.NewApiKeyRepository(global.MySQL)
		apiKeyService := services.NewApiKeyService(apiKeyRepo)
		isExist, err := apiKeyService.CheckApiKey(apiKey)

		if err != nil || !isExist {
			utils.WriteError(c, http.StatusUnauthorized, "Api key Invalid")
			c.Abort()
			return
		}

		c.Next()
	}
}
