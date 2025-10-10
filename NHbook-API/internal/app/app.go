package app

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/initializer"
)

func Run() {
	initializer.InitServer()
	r := initializer.InitRouter()

	r.Run(":3030")
}
