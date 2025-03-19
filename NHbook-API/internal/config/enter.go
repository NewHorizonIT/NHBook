package config

type Config struct {
	Server    `mapstructure:"server"`
	JWTConfig `mapstructure:"JWT"`
	MySQL     `mapstructure:"mysql"`
	Logger    `mapstructure:"logger"`
	Redis     `mapstructure:"redis"`
}
