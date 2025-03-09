package initializer

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/pkg/logger"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/pkg/viper"
)

func InitServer() {
	// Init form package pkg
	viper.InitViper("./configs/")
	logger.InitLogger()

	// Init from package initializer
	InitMySQL()
	InitRedis()
}
