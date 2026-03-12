package config

import (
	"strings"

	"github.com/spf13/viper"
)

type LoadConfig struct {
	Path string
	Name string
	Type string
}

func Load(lc *LoadConfig) (*Config, error) {
	v := viper.New()
	v.SetConfigName(lc.Name)
	v.SetConfigType(lc.Type)
	v.AddConfigPath(lc.Path)

	// Overide with environment variables
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal config into struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Validate config
	if err := ValidateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
