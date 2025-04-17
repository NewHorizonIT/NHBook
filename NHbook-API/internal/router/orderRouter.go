package router

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/wire"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (or *OrderRouter) SetupRouter(r *gin.RouterGroup) {
	orderRouter := r.Group("/orders/")
	orderHandler, _ := wire.InitOrderHandler(global.MySQL, global.Redis)
	{
		// Create new Order
		orderRouter.POST("/", orderHandler.CreateOrder)
	}
}
