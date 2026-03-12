package domain

import "context"

type IUserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	CheckUserExists(ctx context.Context, username string, email string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
