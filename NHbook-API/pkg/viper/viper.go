package viper

import (
	"fmt"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/config"
	"github.com/spf13/viper"
)


func InitViper (path ...string){
	appConfig := config.Config{}
	v := viper.New()
	v.AddConfigPath(path[0])
	v.SetConfigType("yml")
	v.SetConfigName("local.config")


	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Read Config Error %v", err))
	}

	if err := v.Unmarshal(&appConfig); err != nil {
		panic(fmt.Sprintf("Unmarsal Config Error %v", err))
	}

	global.Viper = v
	global.Config = appConfig
} 