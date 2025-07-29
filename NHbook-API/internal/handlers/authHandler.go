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

// @Summary Register a new user
// @Description Register a new user with username, email, password, and role
// @Tags auth
// @Accept json
// @Produce json
// @Param user body request.RegisterRequest true "User registration details"
// @Success 201 {object} utils.ResponseSuccess{data=response.RegisterResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /auth/register [post]
// @Security ApiKeyAuth
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
	c.SetCookie("refresh-token", res.RefreshToken, 604800, "/", "localhost", false, true)

	utils.WriteResponse(c, http.StatusCreated, "Create user success", res, nil)
}

// @Summary User login
// @Description User login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body request.LoginRequest true "User login details"
// @Success 200 {object} utils.ResponseSuccess{data=response.LoginResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /auth/login [post]
// @Security ApiKeyAuth

func (ah *AuthHandler) Login(c *gin.Context) {
	var user request.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.WriteError(c, http.StatusUnauthorized, REQUEST_BODY_INVALID)
		return
	}

	res, err := ah.AuthService.Login(&user)

	if err != nil {
		utils.WriteError(c, http.StatusUnauthorized, LOGIN_UNSUCCESS)
		return
	}
	c.SetCookie("refresh-token", res.RefreshToken, 604800, "/", "localhost", false, true)
	utils.WriteResponse(c, http.StatusOK, LOGIN_SUCCESS, res, nil)

}
func (ah *AuthHandler) Logout(c *gin.Context) {

}

// @Summary Handle refresh token
// @Description Handle refresh token to get new access token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body request.HandleRefreshTokenRequest true "Refresh token details"
// @Success 201 {object} utils.ResponseSuccess{data=response.HandleRefreshTokenResponse}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /auth/refresh-token [post]
// @Security ApiKeyAuth
// @Security ApiKeyAuth
// HandleRefreshToken handles the refresh token request to generate a new access token
// It expects a JSON body with the refresh token and returns a new access token if successful.
func (ah *AuthHandler) HandleRefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh-token")
	if err != nil {
		utils.WriteError(c, http.StatusUnauthorized, REQUEST_BODY_INVALID)
	}
	res, err := ah.AuthService.HandleRefreshToken(refreshToken)

	if err != nil {
		utils.WriteError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.SetCookie("refresh-token", res.RefreshToken, 604800, "/", "localhost", false, true)

	utils.WriteResponse(c, http.StatusCreated, HANDLE_REFRESH_TOKEN_SUCCESS, res, nil)

}
