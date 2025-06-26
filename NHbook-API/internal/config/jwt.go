package config

type JWT struct {
	Secret             string `mapstructure:"secret"`
	Algorithm          string `mapstructure:"algorithm"`
	AccessTokenExpiry  string `mapstructure:"accessTokenExpiry"`
	RefreshTokenExpiry string `mapstructure:"refreshTokenExpiry"`
}
