package app

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/initializer"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/migrations"
)

func Run() {
	initializer.InitServer()
	r := initializer.InitRouter()
	migrations.Migrate(global.MySQL)

	r.Run(":3030")
}
