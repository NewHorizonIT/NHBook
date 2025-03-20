package middlewares

import (
	"net/http"
	"strings"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

var PerfixToken = "Bearer "

func AuthMiddlerware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(global.HEADER_AUTHORIZATION)
		if tokenString == "" {
			utils.WriteError(c, http.StatusForbidden, "Missing token in header")
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, PerfixToken) {
			utils.WriteError(c, http.StatusUnauthorized, "Invalid token format")
		}

		token := strings.TrimPrefix(tokenString, PerfixToken)

		verify, err := utils.VerifyToken(token)

		if err != nil {
			utils.WriteError(c, http.StatusForbidden, "Verify token error")
			c.Abort()
			return
		}

		c.Set("userID", verify.UserID)
		c.Set("email", verify.Email)
		c.Next()
	}
}
