package services

import (
	"errors"
	"fmt"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
)

// Custom Error
var (
	ErrCheckEmail       = errors.New("failed query email")
	ErrEmailExists      = errors.New("email already exists")
	ErrEmailNotExists   = errors.New("email isn't exists")
	ErrHashPasswordFail = errors.New("failed to hash password")
	ErrCreateUserFail   = errors.New("failed to create user")
	ErrTokenGenFail     = errors.New("failed to generate token")
	ErrTokenSaveFail    = errors.New("failed to save token")
	ErrMatchPassword    = errors.New("password don't match")
	ErrVerifyToken      = errors.New("failed to verify token")
	ErrGetToken         = errors.New("failed to get token")
	ErrTokenInvalid     = errors.New("token invalid")
)

// Custom code
var (
	AuthSuccess = "AS1"
	AuthErr     = "AE1"
)

type IAuthService interface {
	Register(username string, email string, password string, roleName string) (*response.RegisterResponse, error)
	Login(*request.LoginRequest) (*response.LoginResponse, error)
	Logout() (map[string]any, error)
	HandleRefreshToken(*request.HandleRefreshTokenRequest) (*response.HandleRefreshTokenResponse, error)
}

type authService struct {
	UserRepo  repositories.IUserRepository
	TokenRepo repositories.ITokenRepository
}

func NewAuthService(ur repositories.IUserRepository, tr repositories.ITokenRepository) IAuthService {
	return &authService{
		UserRepo:  ur,
		TokenRepo: tr,
	}
}

// HandleRefreshToken implements IAuthService.
func (a *authService) HandleRefreshToken(req *request.HandleRefreshTokenRequest) (*response.HandleRefreshTokenResponse, error) {
	// Step 1: Verify token
	claim, err := utils.VerifyToken(req.RefreshToken)

	if err != nil {
		return nil, utils.FormatError(ErrVerifyToken, err)
	}

	// Step 2: Check token
	tokenStore, err := a.TokenRepo.GetTokenByID(claim.UserID)

	if err != nil {
		return nil, utils.FormatError(ErrGetToken, err)
	}
	fmt.Printf("TOKENS: %v \n", tokenStore.Token)
	fmt.Printf("TOKENS BODY: %v \n", req.RefreshToken)

	if tokenStore.Token != req.RefreshToken {
		return nil, utils.FormatError(ErrTokenInvalid, ErrTokenInvalid)
	}

	// Step 3: Generate Access token and Refresh token

	accessToken, refreshToken, err := utils.CreateTokenPair(claim.UserID, claim.Email)

	if err != nil {
		return nil, utils.FormatError(ErrCreateUserFail, err)
	}

	// Step 4: Update token into table tokens
	if err := a.TokenRepo.UpdateTokenByID(claim.UserID, refreshToken); err != nil {
		return nil, utils.FormatError(ErrTokenSaveFail, err)
	}

	// Step 5: Create response

	res := &response.HandleRefreshTokenResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	return res, nil
}

// Login implements IAuthService.
func (a *authService) Login(user *request.LoginRequest) (*response.LoginResponse, error) {
	// Step 1: Check email exist
	foundUser, err := a.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, utils.FormatError(ErrCheckEmail, err)
	}

	if foundUser == nil {
		return nil, utils.FormatError(ErrEmailNotExists)
	}

	// Step 2: Compare password
	isMatch := utils.CompareHashPassword(user.Password, foundUser.Password)

	if !isMatch {
		return nil, utils.FormatError(ErrMatchPassword)
	}

	// Step 3: Create new AccessToken and RefreshToken
	accessToken, refreshToken, err := utils.CreateTokenPair(foundUser.ID, foundUser.Email)

	if err != nil {
		return nil, utils.FormatError(ErrTokenGenFail, err)
	}

	// Step 4: Save token into table tokens

	if err := a.TokenRepo.UpdateTokenByID(foundUser.ID, refreshToken); err != nil {
		return nil, utils.FormatError(ErrTokenSaveFail, err)
	}

	// Step 5: Create Response

	res := &response.LoginResponse{
		Code:         AuthSuccess,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	res.User.UserID = foundUser.ID
	res.User.Email = foundUser.Email
	return res, nil
}

// Logout implements IAuthService.
func (a *authService) Logout() (map[string]any, error) {
	return nil, nil
}

// Register implements IAuthService.
func (a *authService) Register(userName string, email string, password string, roleName string) (*response.RegisterResponse, error) {
	// Step1 1: Check user exits
	holderEmail, err := a.UserRepo.IsEmailExist(email)
	if err != nil {
		return nil, fmt.Errorf("failed check email %w", err)
	}

	if holderEmail {
		return nil, ErrEmailExists
	}

	// Step 2: Create user
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, ErrHashPasswordFail
	}

	newUser, err := a.UserRepo.CreateUser(userName, email, hashPassword, roleName)

	if err != nil {
		return nil, fmt.Errorf("%s: %w ", ErrCreateUserFail, err)
	}

	// Step 3: Create accessToken  and refreshToken
	accessToken, refreshToken, err := utils.CreateTokenPair(newUser.ID, newUser.Email)

	if err != nil {
		return nil, fmt.Errorf("%s: %w ", ErrTokenGenFail, err)
	}

	// Step 4: Save token into table tokens

	err = a.TokenRepo.CreateToken(newUser.ID, refreshToken)

	if err != nil {
		return nil, fmt.Errorf("%s: %w ", ErrTokenSaveFail, err)
	}

	// Step 5: Create response

	res := &response.RegisterResponse{
		Code:         AuthSuccess,
		Accesstoken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Step 6: Create Role User

	res.User.Email = newUser.Email
	res.User.UserID = newUser.ID

	return res, nil

}
