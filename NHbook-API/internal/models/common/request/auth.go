package request

type RegisterRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HandleRefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
