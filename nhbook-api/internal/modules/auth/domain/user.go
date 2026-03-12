package domain

import (
	"github.com/NewHorizonIT/nhbook-api/pkg/auth"
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Username string
	Email    Email
	Password string
	Role     Role
	Status   Status
}

// NewUser tạo user mới với password đã được hash
func NewUser(username string, email string, password string, role string, status string) (*User, error) {
	// 1. Validate và tạo Email value object
	emailVal, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	// 2. Validate password (có thể thêm các rule phức tạp hơn)
	if len(password) < 8 {
		return nil, ErrPasswordTooWeak
	}

	// 3. Hash password
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 4. Tạo user entity
	return &User{
		ID:       uuid.NewString(),
		Username: username,
		Email:    emailVal,
		Password: hashedPassword,
		Role:     NewRole(role),
		Status:   NewStatus(status),
	}, nil
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) GetID() string {
	return user.ID
}
