package app

import (
	"fmt"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/initializer"
	"github.com/gin-gonic/gin"
)

func Run() {
	initializer.InitServer()
	configServer := global.Config.Server
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	r.Run(fmt.Sprintf(":%d", configServer.Port))
}
