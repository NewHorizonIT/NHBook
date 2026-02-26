package usecase

import (
	"context"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/config"
	"github.com/NewHorizonIT/nhbook-api/pkg/auth"
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
	// 1. Kiểm tra user đã tồn tại chưa
	exists, err := uc.repo.CheckUserExists(ctx, dto.Username, dto.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrUsernameAlreadyUsed
	}

	// 2. Tạo domain entity
	user, err := domain.NewUser(dto.Username, dto.Email, dto.Password, dto.Role, dto.Status)
	if err != nil {
		return nil, err
	}

	// 3. Lưu vào database
	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// 5. Generate token
	accessToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.AccessExpirationMinutes))
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateJWT(user.GetID(), user.Role.String(), uc.jwtConf.Secret, time.Duration(uc.jwtConf.RefreshExpirationMinutes))
	if err != nil {
		return nil, err
	}

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
