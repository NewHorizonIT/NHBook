package handlers

import (
	"net/http"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

// Message
const (
	REQUEST_BODY_INVALID           = "Request body invalid"
	LOGIN_UNSUCCESS                = "Login unsuccess"
	LOGIN_SUCCESS                  = "Login success"
	HANDLE_REFRESH_TOKEN_UNSUCCESS = "Handle RefreshToken unsuccess"
	HANDLE_REFRESH_TOKEN_SUCCESS   = "Handle RefreshToken success"
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

	var user request.RegisterRequest
	err := c.ShouldBindJSON(&user)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, REQUEST_BODY_INVALID)
		return
	}

	res, err := ah.AuthService.Register(user.UserName, user.Email, user.Password, user.Role)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteResponse(c, http.StatusCreated, "Create user success", res, nil)
}
func (ah *AuthHandler) Login(c *gin.Context) {
	var user request.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.WriteError(c, http.StatusBadRequest, REQUEST_BODY_INVALID)
		return
	}

	res, err := ah.AuthService.Login(&user)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, LOGIN_UNSUCCESS)
		return
	}
	utils.WriteResponse(c, http.StatusOK, LOGIN_SUCCESS, res, nil)

}
func (ah *AuthHandler) Logout(c *gin.Context) {

}
func (ah *AuthHandler) HandleRefreshToken(c *gin.Context) {
	var body = request.HandleRefreshTokenRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.WriteError(c, http.StatusBadRequest, REQUEST_BODY_INVALID)
		return
	}

	res, err := ah.AuthService.HandleRefreshToken(&body)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteResponse(c, http.StatusCreated, HANDLE_REFRESH_TOKEN_SUCCESS, res, nil)

}
