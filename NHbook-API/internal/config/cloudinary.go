package config

type Cloudinary struct {
	CloudName string `mapstructure:"cloudName"`
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
}
