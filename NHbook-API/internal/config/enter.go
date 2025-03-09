package config

type Config struct {
	Server `mapstructure:"server"`
	MySQL  `mapstructure:"mysql"`
	Logger `mapstructure:"logger"`
	Redis  `mapstructure:"redis"`
}
