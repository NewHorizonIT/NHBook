package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/db"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IAuthService interface {
	Register(ctx *gin.Context, userDataRegister *request.Register) (*response.AuthData, error)
	Login(userDataLogin *request.Login) (*response.AuthData, error)
	Logout() (map[string]any, error)
	HandleRefreshToken(refreshToken string, refreshTokenInStore string, user utils.Claim) (*response.RefreshTokenData, error)
	GetInfoUser(id string) (*response.AuthData, error)
}

type authService struct {
	UserRepo repositories.IUserRepository
}

func NewAuthService(ur repositories.IUserRepository) IAuthService {
	return &authService{
		UserRepo: ur,
	}
}

// HandleRefreshToken implements IAuthService.
func (a *authService) HandleRefreshToken(refreshToken string, refreshTokenInStore string, user utils.Claim) (*response.RefreshTokenData, error) {

	// Step 1: Check token
	if refreshToken != refreshTokenInStore {
		return nil, fmt.Errorf("compare token error: %w", errors.New("token no macth"))
	}
	// Step 2: Generate Access token and Refresh token

	newAccessToken, newRefreshToken, err := utils.CreateTokenPair(user.UserID, user.Email)

	if err != nil {
		return nil, fmt.Errorf("create token error: %w", err)
	}

	// Step 3: Create data
	data := &response.RefreshTokenData{
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}
	return data, nil
}

// Login implements IAuthService.
func (a *authService) Login(userDataLogin *request.Login) (*response.AuthData, error) {
	// Step 1: Check email exist
	ctx := context.Background()
	foundUser, err := a.UserRepo.GetUserByEmail(ctx, userDataLogin.Email)
	if err != nil {
		return nil, fmt.Errorf("get user by email error: %w", err)
	}

	// Step 2: Compare password
	isMatch := utils.CompareHashPassword(userDataLogin.Password, foundUser.PasswordHash)

	if !isMatch {
		return nil, fmt.Errorf("compare password error: %w", errors.New("password no macth"))
	}

	// Step 3: Create new AccessToken and RefreshToken
	accessToken, refreshToken, err := utils.CreateTokenPair(foundUser.ID, foundUser.Email)

	if err != nil {
		return nil, fmt.Errorf("create token error: %w", err)
	}

	// Step 5: Create Login Data
	dataUser := response.UserData{
		ID:        foundUser.ID,
		Username:  foundUser.Username,
		Email:     foundUser.Email,
		Phone:     foundUser.Phone.String,
		CreatedAt: foundUser.CreatedAt.Time,
	}
	dataLogin := &response.AuthData{
		User:         dataUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return dataLogin, nil
}

// Logout implements IAuthService.
func (a *authService) Logout() (map[string]any, error) {
	return nil, nil
}

// Register implements IAuthService.
func (a *authService) Register(ctx *gin.Context, userDataRegister *request.Register) (*response.AuthData, error) {
	// Step1 1: Check user exits
	_, err := a.UserRepo.GetUserByEmail(ctx, userDataRegister.Email)
	if err != nil || errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("check email error: %w", err)
	}

	// Step 2: Hash password
	hashPassword, err := utils.HashPassword(userDataRegister.Password)
	if err != nil {
		return nil, fmt.Errorf("hash passowrd error: %w", err)
	}

	// Step 3: Create userParams
	id := uuid.New()
	createUserData := db.CreateUserParams{
		ID:           id.String(),
		Username:     userDataRegister.UserName,
		Email:        userDataRegister.Email,
		PasswordHash: hashPassword,
		Phone:        utils.ToNullString(userDataRegister.PhoneNumber),
	}
	newUser, err := a.UserRepo.CreateUser(ctx, createUserData)

	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	// Step 3: Create accessToken  and refreshToken
	accessToken, refreshToken, err := utils.CreateTokenPair(newUser.ID, newUser.Email)

	if err != nil {
		return nil, fmt.Errorf("create token error: %w", err)
	}

	// Step 4: Create Auth data
	dataUser := response.UserData{
		ID:          newUser.ID,
		Email:       newUser.Email,
		Phone:       newUser.Phone.String,
		CreatedAt:   newUser.CreatedAt.Time,
		IsAdmin:     newUser.IsAdmin.Bool,
		DisplayName: newUser.DisplayName.String,
	}
	dataRegister := &response.AuthData{
		User:         dataUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return dataRegister, nil

}

// GetInfoUser implements IAuthService.
func (a *authService) GetInfoUser(id string) (*response.AuthData, error) {
	panic("Not implement")
}
