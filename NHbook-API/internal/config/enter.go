package config

type Config struct {
	Server     `mapstructure:"server"`
	JWT        `mapstructure:"JWT"`
	MySQL      `mapstructure:"mysql"`
	Logger     `mapstructure:"logger"`
	Redis      `mapstructure:"redis"`
	Cloudinary `mapstructure:"cloudinary"`
}
