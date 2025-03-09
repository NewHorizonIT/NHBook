package config

type MySQL struct {
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	Name               string `mapstructure:"name"`
	MaxIdleConnect     int    `mapstructure:"maxIdleConnect"`
	MaxOpenConnect     int    `mapstructure:"maxOpenConnect"`
	MaxConnectTimeLife int    `mapstructure:"maxConnectTimeLife"`
}
