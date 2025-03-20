package response

type RegisterResponse struct {
	Code string `json:"code"`
	User struct {
		UserID string `json:"userID"`
		Email  string `json:"email"`
	} `json:"user"`
	Accesstoken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	Code string `json:"code"`
	User struct {
		UserID string `json:"userID"`
		Email  string `json:"email"`
	} `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type HandleRefreshTokenResponse struct {
	Code         string `json:"code"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
