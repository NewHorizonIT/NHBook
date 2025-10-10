package response

import (
	"time"
)

// Define Base Respone
type MetaResponse struct {
	Total int32 `json:"total"`
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}
type BaseResponse struct {
	Code      string `json:"code"`
	IsSuccess bool   `json:"is_success"`
	Error     string `json:"error"`
}

type ListReponse[T any] struct {
	BaseResponse
	Data []T `json:"data"`
	MetaResponse
}

// Define Authenticate Response

type UserData struct {
	ID          string
	Username    string
	Email       string
	DisplayName string
	Phone       string
	IsAdmin     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AuthData struct {
	User         UserData `json:"user"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
}
type Authentication struct {
	BaseResponse
	Data AuthData `json:"data"`
}

type RefreshTokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type HandleRefreshToken struct {
	BaseResponse
	Data RefreshTokenData `json:"data"`
}
