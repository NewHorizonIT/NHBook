package services

import "github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"

type IAuthService interface {
	Register(username string, email string, password string) (map[string]any, error)
	Login(email string, password string) (map[string]any, error)
	Logout() (map[string]any, error)
	HandleRefreshToken() (map[string]any, error)
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
func (a *authService) HandleRefreshToken() (map[string]any, error) {
	return nil, nil
}

// Login implements IAuthService.
func (a *authService) Login(email string, password string) (map[string]any, error) {
	return nil, nil
}

// Logout implements IAuthService.
func (a *authService) Logout() (map[string]any, error) {
	return nil, nil
}

// Register implements IAuthService.
func (a *authService) Register(username string, email string, password string) (map[string]any, error) {
	return nil, nil
}
