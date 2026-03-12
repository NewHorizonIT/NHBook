package infrastructure

import (
	"context"
	"database/sql"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/auth/infrastructure/sqlc"
	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func NewUserRepository(db *sql.DB) domain.IUserRepository {
	return &UserRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

// CreateUser chuyển đổi domain.User sang sqlc params và lưu vào DB
func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	userID, err := uuid.Parse(user.ID)
	if err != nil {
		return errs.Wrap(err, "UserRepository.CreateUser", domain.ErrCodeInternalServerError, "invalid user ID format")
	}

	params := sqlc.CreateUserParams{
		ID:       userID,
		Email:    user.Email.String(),
		Password: user.Password, // Password đã được hash ở domain layer
		Username: user.Username,
		Role:     sql.NullString{String: user.Role.String(), Valid: true},
		Status:   sql.NullString{String: user.Status.String(), Valid: true},
	}

	err = r.q.CreateUser(ctx, params)
	if err != nil {
		return errs.Wrap(err, "UserRepository.CreateUser", domain.ErrCodeInternalServerError, "failed to create user")
	}

	return nil
}

func (r *UserRepository) CheckUserExists(ctx context.Context, username string, email string) (bool, error) {
	params := sqlc.UserExistsParams{
		Username: username,
		Email:    email,
	}
	exists, err := r.q.UserExists(ctx, params)
	if err != nil {
		return false, errs.Wrap(err, "UserRepository.CheckUserExists", domain.ErrCodeInternalServerError, "failed to check user exists")
	}

	return exists, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	dbUser, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, errs.Wrap(err, "UserRepository.GetUserByEmail", domain.ErrCodeInternalServerError, "failed to get user by email")
	}

	// Chuyển đổi từ DB model sang domain entity
	user := &domain.User{
		ID:       dbUser.ID.String(),
		Username: dbUser.Username,
		Email:    domain.Email(dbUser.Email),
		Password: dbUser.Password,
		Role:     domain.NewRole(dbUser.Role.String),
		Status:   domain.NewStatus(dbUser.Status.String),
	}

	return user, nil
}
