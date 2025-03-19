package config

type JWTConfig struct {
	Secret             string `mapstructure:"secret"`
	Algorithm          string `mapstructure:"algorithm"`
	AccessTokenExpiry  string `mapstructure:"access_token_expiry"`
	RefreshTokenExpiry string `mapstructure:"refresh_token_expiry"`
}
