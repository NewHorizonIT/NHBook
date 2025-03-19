package handlers

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService services.IAuthService
}

func NewAuthHandler(as services.IAuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: as,
	}
}

func (ah *AuthHandler) Register(c *gin.Context) {

}
func (ah *AuthHandler) Login(c *gin.Context) {

}
func (ah *AuthHandler) Logout(c *gin.Context) {

}
func (ah *AuthHandler) HandleRefreshToken(c *gin.Context) {

}
