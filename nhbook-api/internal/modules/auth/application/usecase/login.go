package usecase

import (
	"context"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/NewHorizonIT/nhbook-api/pkg/auth"
	"github.com/NewHorizonIT/nhbook-api/pkg/logger"
	"go.uber.org/zap"
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
	log := logger.FromContext(ctx)
	
	// Step 1: Check user is exists
	log.Debug("Login: Fetching user by email", zap.String("email", dto.Email))
	user, err := uc.repo.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		log.Warn("Login: User not found", zap.String("email", dto.Email), zap.Error(err))
		return nil, err
	}

	// Step 2: Validate password
	if err := auth.ComparePassword(user.GetPassword(), dto.Password); err != nil {
		log.Warn("Login: Invalid password", zap.String("email", dto.Email))
		return nil, domain.ErrInvalidCredentials
	}

	// Step 3: Generate Pair token
	accessToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.AccessExpirationMinutes))
	if err != nil {
		log.Error("Login: Failed to generate access token", zap.Error(err))
		return nil, err
	}

	refreshToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.RefreshExpirationMinutes))
	if err != nil {
		log.Error("Login: Failed to generate refresh token", zap.Error(err))
		return nil, err
	}

	log.Info("Login: User logged in successfully", zap.String("user_id", user.GetID()))

	return &application.LoginResponseDTO{
		ID:           user.GetID(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
