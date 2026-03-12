package domain

import "errors"

// Bussiness Error Definitions
var (
	ErrUserNotFound        = errors.New("User not found")
	ErrEmailIsRequired     = errors.New("Email is required")
	ErrInvalidEmailFormat  = errors.New("Invalid email format")
	ErrPasswordTooWeak     = errors.New("Password is too weak")
	ErrUsernameAlreadyUsed = errors.New("Username is already used")
	ErrInvalidStatus       = errors.New("Invalid status value")
	ErrInvalidCredentials  = errors.New("Invalid credentials")
)

// Error codes
const (
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
)
