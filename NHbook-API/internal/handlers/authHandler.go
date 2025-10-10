package handlers

import (
	"fmt"
	"net/http"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Message
const (
	REGISTER_FAIL                  = "register fail"
	REGISTER_SUCCESS               = "register success"
	REQUEST_BODY_INVALID           = "Request body invalid"
	LOGIN_UNSUCCESS                = "Login unsuccess"
	LOGIN_SUCCESS                  = "Login success"
	HANDLE_REFRESH_TOKEN_UNSUCCESS = "handle refreshToken unsuccess"
	HANDLE_REFRESH_TOKEN_SUCCESS   = "Handle RefreshToken success"
	UNAUTHENTICATION               = "UNAUTHENTIACTION ERROR"
	SOMTHING_WENT_WRONG            = "SOMETHING WENT WRONG"
)

type AuthHandler struct {
	AuthService  services.IAuthService
	CacheService services.ICache
}

func NewAuthHandler(as services.IAuthService, cs services.ICache) *AuthHandler {
	return &AuthHandler{
		AuthService:  as,
		CacheService: cs,
	}
}

// @Summary Register a new user
// @Description Register a new user with username, email, password, and phone number
// @Tags auth
// @Accept json
// @Produce json
// @Param user body request.Register true "User registration details"
// @Success 201 {object} utils.ResponseSuccess{data=response.AuthData}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /auth/register [post]
// @Security ApiKeyAuth
func (ah *AuthHandler) Register(ctx *gin.Context) {
	var user request.Register
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		utils.WriteError(ctx, http.StatusBadRequest, REQUEST_BODY_INVALID)
		return
	}

	res, err := ah.AuthService.Register(ctx, &user)
	if err != nil {
		global.Logger.Error("register faild", zap.String("error", err.Error()))
		utils.WriteError(ctx, http.StatusBadRequest, REGISTER_FAIL)
		return
	}
	ctx.SetCookie("refresh-token", res.RefreshToken, 604800, "/", "localhost", false, true)
	keyCache := fmt.Sprintf("refresh:%v", res.User.ID)
	ah.CacheService.Set(ctx, keyCache, res.RefreshToken, 7*24*60*60)

	utils.WriteResponse(ctx, http.StatusCreated, REGISTER_SUCCESS, res, nil)
}

// @Summary User login
// @Description User login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body request.Login true "User login details"
// @Success 200 {object} utils.ResponseSuccess{data=response.AuthData}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /auth/login [post]
// @Security ApiKeyAuth
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var user request.Login

	if err := ctx.ShouldBindJSON(&user); err != nil {
		global.Logger.Error("Bind body error", zap.String("error", err.Error()))
		utils.WriteError(ctx, http.StatusUnauthorized, REQUEST_BODY_INVALID)
		return
	}

	res, err := ah.AuthService.Login(&user)
	if err != nil {
		global.Logger.Error("Login error", zap.String("error", err.Error()))
		utils.WriteError(ctx, http.StatusUnauthorized, LOGIN_UNSUCCESS)
		return
	}
	ctx.SetCookie("refresh-token", res.RefreshToken, 604800, "/", "localhost", false, true)
	keyCache := fmt.Sprintf("refresh:%v", res.User.ID)
	ah.CacheService.Set(ctx, keyCache, res.RefreshToken, 7*24*60*60)
	utils.WriteResponse(ctx, http.StatusOK, LOGIN_SUCCESS, res, nil)
}

func (ah *AuthHandler) Logout(c *gin.Context) {
}

// @Summary Handle refresh token
// @Description Handle refresh token to get new access token
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} utils.ResponseSuccess{data=response.RefreshTokenData}
// @Failure 400 {object} utils.ResponseError{message=string}
// @Router /auth/refresh-token [post]
// @Security ApiKeyAuth
// HandleRefreshToken handles the refresh token request to generate a new access token
// It expects a JSON body with the refresh token and returns a new access token if successful.
func (ah *AuthHandler) HandleRefreshToken(ctx *gin.Context) {
	// Step 1: Get refresh oken from cookie
	refreshToken, err := ctx.Cookie("refresh-token")
	if err != nil {
		global.Logger.Error("Bind body error", zap.String("error", err.Error()))
		utils.WriteError(ctx, http.StatusUnauthorized, REQUEST_BODY_INVALID)
	}

	// Step 2: verify refresh token
	data, err := utils.VerifyToken(refreshToken)

	if err != nil {
		global.Logger.Error("Verify token error", zap.String("error", err.Error()))
		utils.WriteError(ctx, http.StatusUnauthorized, HANDLE_REFRESH_TOKEN_UNSUCCESS)
		return
	}
	// Step 3: Get token from redis
	keyCache := fmt.Sprintf("refresh:%v", data.UserID)
	refreshTokenInStore, err := ah.CacheService.Get(ctx, keyCache)
	if err != nil || refreshTokenInStore != refreshToken {
		global.Logger.Error("Get token from redis error", zap.String("error", err.Error()))
	}

	// Step 4: Call service to handle refresh token
	res, err := ah.AuthService.HandleRefreshToken(refreshToken, refreshTokenInStore, *data)
	if err != nil {
		global.Logger.Error("Handle refresh token error", zap.String("error", err.Error()))
		utils.WriteError(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.SetCookie("refresh-token", res.RefreshToken, 604800, "/", "localhost", false, true)

	// Step 5: Update token in redis
	ah.CacheService.Set(ctx, keyCache, res.RefreshToken, 7*24*60*60)
	utils.WriteResponse(ctx, http.StatusCreated, HANDLE_REFRESH_TOKEN_SUCCESS, res, nil)
}

func (ah *AuthHandler) GetInfoUser(c *gin.Context) {
	userId := c.GetString("userID")

	if userId == "" {
		utils.WriteError(c, http.StatusForbidden, UNAUTHENTICATION)
		return
	}
}
