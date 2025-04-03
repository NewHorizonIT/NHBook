package router

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/wire"
	"github.com/gin-gonic/gin"
)

type BookRouter struct {
}

func (br *BookRouter) SetUpBookRouter(r *gin.RouterGroup) {
	bookHandler, _ := wire.InitBookHandler(global.MySQL)
	bookRouter := r.Group("/books")
	{
		bookRouter.POST("/", bookHandler.CreateBook)
		bookRouter.GET("/", bookHandler.GetListBook)
	}

}
