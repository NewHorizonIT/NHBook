package presentation

import (
	"net/http"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/domain"
	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/NewHorizonIT/nhbook-api/pkg/logger"
	"github.com/NewHorizonIT/nhbook-api/pkg/response"
	"github.com/NewHorizonIT/nhbook-api/pkg/validator"
	"github.com/gin-gonic/gin"
	validatorpkg "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// AuthHandler chứa tất cả dependencies cần thiết
type AuthHandler struct {
	createUserUsecase *usecase.CreateUserUsecase
	loginUsecase      *usecase.LoginUsecase
	validator         *validatorpkg.Validate
}

// NewAuthHandler khởi tạo handler với dependency injection
func NewAuthHandler(
	createUserUsecase *usecase.CreateUserUsecase,
	loginUsecase *usecase.LoginUsecase,
) *AuthHandler {
	return &AuthHandler{
		createUserUsecase: createUserUsecase,
		loginUsecase:      loginUsecase,
		validator:         validatorpkg.New(),
	}
}

// SignUp godoc
// @Summary      Register new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body application.CreateUserDTO true "User registration data"
// @Success      201 {object} application.CreateUserResponseDTO
// @Router       /auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContext(ctx)

	var req application.CreateUserDTO

	// 1. Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warn("SignUp: Invalid request body", zap.Error(err))
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// 2. Validate
	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		log.Warn("SignUp: Validation failed", zap.Any("errors", validationErrors))
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	log.Info("SignUp: Creating new user", zap.String("email", req.Email), zap.String("username", req.Username))

	// 3. Execute usecase
	result, err := h.createUserUsecase.Execute(ctx, &req)
	if err != nil {
		log.Error("SignUp: Failed to create user", zap.Error(err))
		h.handleError(c, err)
		return
	}

	// 4. Set refresh token in HttpOnly cookie
	c.SetCookie("refresh_token", result.RefreshToken, 3600*24*7, "/", "", false, true)

	log.Info("SignUp: User created successfully", zap.String("user_id", result.ID))

	// 5. Success response
	response.WriteSuccessResponse(c, http.StatusCreated, result)
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body application.LoginRequestDTO true "Login credentials"
// @Success      200 {object} application.LoginResponseDTO
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContext(ctx)

	var req application.LoginRequestDTO

	// 1. Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warn("Login: Invalid request body", zap.Error(err))
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// 2. Validate
	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		log.Warn("Login: Validation failed", zap.Any("errors", validationErrors))
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	log.Info("Login: User attempting login", zap.String("email", req.Email))

	// 3. Execute usecase
	result, err := h.loginUsecase.Execute(ctx, &req)
	if err != nil {
		log.Warn("Login: Failed login attempt", zap.String("email", req.Email), zap.Error(err))
		h.handleError(c, err)
		return
	}

	// 4. Set refresh token in HttpOnly cookie
	c.SetCookie("refresh_token", result.RefreshToken, 3600*24*7, "/", "", false, true)

	log.Info("Login: User logged in successfully", zap.String("user_id", result.ID))

	// 5. Success response
	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// handleError xử lý lỗi từ domain/application layer
func (h *AuthHandler) handleError(c *gin.Context, err error) {
	// Kiểm tra custom domain errors
	switch err {
	case domain.ErrUserNotFound, domain.ErrInvalidCredentials:
		response.WriteErrorResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
	case domain.ErrEmailIsRequired, domain.ErrInvalidEmailFormat:
		response.WriteErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
	case domain.ErrUsernameAlreadyUsed:
		response.WriteErrorResponse(c, http.StatusConflict, err.Error(), nil)
	default:
		// Kiểm tra wrapped errors
		if wrappedErr, ok := err.(*errs.AppError); ok {
			response.WriteErrorResponse(c, http.StatusInternalServerError, wrappedErr.Message, wrappedErr.Code)
		} else {
			response.WriteErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
		}
	}
}
