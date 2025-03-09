package config

type Server struct {
	AppName string `mapstructure:"appName"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Env     string `mapstructure:"env"`
}
