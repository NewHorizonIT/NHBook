package config

import "github.com/NewHorizonIT/nhbook-api/pkg/logger"

type Config struct {
	App   AppConfig           `mapstructure:"app" validate:"required"`
	Log   logger.ConfigLogger `mapstructure:"log" validate:"required"`
	Db    DbConfig            `mapstructure:"db" validate:"required"`
	Redis RedisConfig         `mapstructure:"redis" validate:"required"`
	JWT   JWTConfig           `mapstructure:"jwt" validate:"required"`
}

type AppConfig struct {
	Host            string `mapstructure:"host" validate:"required"`
	Port            int    `mapstructure:"port" validate:"required"`
	Name            string `mapstructure:"name" validate:"required"`
	DevelopmentMode bool   `mapstructure:"development_mode" validate:"required"`
}

type DbConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
	MaxIdle  int    `mapstructure:"max_idle" validate:"required"`
	MaxOpen  int    `mapstructure:"max_open" validate:"required"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Password string `mapstructure:"password" default:""`
	Prefix   string `mapstructure:"prefix" validate:"required"`
}

type JWTConfig struct {
	Secret                   string `mapstructure:"secret" validate:"required"`
	AccessExpirationMinutes  int    `mapstructure:"exp_access_minutes" validate:"required"`
	RefreshExpirationMinutes int    `mapstructure:"exp_refresh_minutes" validate:"required"`
}
