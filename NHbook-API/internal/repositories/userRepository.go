package repositories

import (
	"context"
	"database/sql"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/db"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	UpdateUser(ctx context.Context, id string, payload map[string]any) (db.User, error)
	GetAllUser(ctx context.Context) ([]db.User, error)
	GetUserByID(ctx context.Context, id string) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	q *db.Queries
}

func NewUserRepository(dbCon *sql.DB) IUserRepository {
	return &userRepository{
		q: db.New(dbCon),
	}
}

// CreateUser implements IUserRepository.
func (u *userRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {

	// Step 1: Create user
	_, err := u.q.CreateUser(ctx, arg)
	if err != nil {
		return db.User{}, err
	}
	// Step 2: Get user by ID
	res, err := u.q.GetUserByID(ctx, arg.ID)
	if err != nil {
		return db.User{}, err
	}

	// Step 3: convert to db.User and return
	user := db.User{
		ID:           res.ID,
		Username:     res.Username,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Phone:        res.Phone,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}
	return user, err
}

// GetAllUser implements IUserRepository.
func (u *userRepository) GetAllUser(ctx context.Context) ([]db.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	// Step 1: Get User by email
	user, err := u.q.GetUserByEmail(ctx, email)

	if err != nil {
		return db.User{}, nil
	}

	// Step 2: Convert to db.user
	convetUser := db.User{
		ID:           user.ID,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Phone:        user.Phone,
		CreatedAt:    user.CreatedAt,
	}

	return convetUser, nil
}

// GetUserByID implements IUserRepository.
func (u *userRepository) GetUserByID(ctx context.Context, id string) (db.User, error) {
	// Step 1: Call method getUserByID
	user, err := u.q.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, err
	}

	// Step 2: Convert to db.User
	convertUser := db.User{
		ID:           user.ID,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Phone:        user.Phone,
		CreatedAt:    user.CreatedAt,
	}

	return convertUser, nil
}

// UpdateUser implements IUserRepository.
func (u *userRepository) UpdateUser(ctx context.Context, id string, payload map[string]any) (db.User, error) {
	panic("unimplemented")
}

func (u *userRepository) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}
