package middlewares

import (
	"net/http"
	"strings"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/pkg/utils"
	"github.com/gin-gonic/gin"
)

var PerfixToken = "Bearer "

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(global.HEADER_AUTHORIZATION)
		if tokenString == "" {
			utils.WriteError(c, http.StatusUnauthorized, "Missing token in header")
			c.Abort()
			return
		}
		if !strings.HasPrefix(tokenString, PerfixToken) {
			utils.WriteError(c, http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}

		token := strings.TrimPrefix(tokenString, PerfixToken)

		verify, err := utils.VerifyToken(token)

		if err != nil {
			utils.WriteError(c, http.StatusUnauthorized, "Verify token error")
			c.Abort()
			return
		}

		c.Set("userID", verify.UserID)
		c.Set("email", verify.Email)
		c.Set("isAdmin", verify.IsAdmin)
		c.Next()
	}
}
