package router

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/wire"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
}

func (ct *CategoryRouter) SetUpRouter(r *gin.RouterGroup) {
	categoryHandler, _ := wire.InitCategoryHanler(global.MySQL)
	categoryRouter := r.Group("/categories")
	{
		categoryRouter.GET("/", categoryHandler.GetListCategory)
		categoryRouter.GET("/status/:status", categoryHandler.GetListCategoryByStatus)
		categoryRouter.POST("/", categoryHandler.CreateCategory)
		categoryRouter.PUT("/", categoryHandler.UpdateCategory)
	}
}
