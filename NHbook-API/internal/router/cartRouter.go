package router

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/middlewares"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/wire"
	"github.com/gin-gonic/gin"
)

type CartRouter struct {
}

func (cr *CartRouter) SetUpRouter(r *gin.RouterGroup) {
	cartHandler, _ := wire.IniCartHandler(global.Redis, global.MySQL)
	cartRouter := r.Group("/carts/")
	{
		cartRouter.Use(middlewares.AuthMiddlerware())
		// Add item to cart
		cartRouter.POST("/", cartHandler.AddItemToCart)
		// Get cart
		cartRouter.GET("/", cartHandler.GetCart)
		// Remove item in cart
		cartRouter.DELETE("/:bookID", cartHandler.RemoveItemInCart)
		// Remove all item in cart
		cartRouter.DELETE("/", cartHandler.RemoveAllItemToCart)
	}
}
