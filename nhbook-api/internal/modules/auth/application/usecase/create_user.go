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

type CreateUserUsecase struct {
	repo    domain.IUserRepository
	jwtConf config.JWTConfig
}

func NewCreateUserUsecase(repo domain.IUserRepository, jwtConf config.JWTConfig) *CreateUserUsecase {
	return &CreateUserUsecase{
		repo:    repo,
		jwtConf: jwtConf,
	}
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, dto *application.CreateUserDTO) (*application.CreateUserResponseDTO, error) {
	log := logger.FromContext(ctx)
	
	// 1. Kiểm tra user đã tồn tại chưa
	log.Debug("CreateUser: Checking if user exists", zap.String("username", dto.Username), zap.String("email", dto.Email))
	exists, err := uc.repo.CheckUserExists(ctx, dto.Username, dto.Email)
	if err != nil {
		log.Error("CreateUser: Failed to check user existence", zap.Error(err))
		return nil, err
	}
	if exists {
		log.Warn("CreateUser: User already exists", zap.String("username", dto.Username))
		return nil, domain.ErrUsernameAlreadyUsed
	}

	// 2. Tạo domain entity
	user, err := domain.NewUser(dto.Username, dto.Email, dto.Password, dto.Role, dto.Status)
	if err != nil {
		log.Error("CreateUser: Failed to create domain entity", zap.Error(err))
		return nil, err
	}

	// 3. Lưu vào database
	log.Debug("CreateUser: Saving user to database", zap.String("user_id", user.GetID()))
	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		log.Error("CreateUser: Failed to save user to database", zap.Error(err))
		return nil, err
	}

	// 5. Generate token
	accessToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.AccessExpirationMinutes))
	if err != nil {
		log.Error("CreateUser: Failed to generate access token", zap.Error(err))
		return nil, err
	}

	refreshToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.RefreshExpirationMinutes))
	if err != nil {
		log.Error("CreateUser: Failed to generate refresh token", zap.Error(err))
		return nil, err
	}

	log.Info("CreateUser: User created successfully", zap.String("user_id", user.GetID()))

	// 6. Trả về response
	response := &application.CreateUserResponseDTO{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email.String(),
		Role:         user.Role.String(),
		Status:       user.Status.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, nil
}
