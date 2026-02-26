package usecase

import (
	"context"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/NewHorizonIT/nhbook-api/pkg/auth"
)

type LoginUsecase struct {
	repo    domain.IUserRepository
	jwtConf config.JWTConfig
}

func NewLoginUsecase(repo domain.IUserRepository, jwtConf config.JWTConfig) *LoginUsecase {
	return &LoginUsecase{
		repo:    repo,
		jwtConf: jwtConf,
	}
}

func (uc *LoginUsecase) Execute(ctx context.Context, dto *application.LoginRequestDTO) (*application.LoginResponseDTO, error) {
	// Step 1: Check user is exists
	user, err := uc.repo.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	// Step 2: Validate password
	if err := auth.ComparePassword(user.GetPassword(), dto.Password); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Step 3: Generate Pair token
	accessToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.AccessExpirationMinutes))
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.RefreshExpirationMinutes))
	if err != nil {
		return nil, err
	}

	return &application.LoginResponseDTO{
		ID:           user.GetID(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
