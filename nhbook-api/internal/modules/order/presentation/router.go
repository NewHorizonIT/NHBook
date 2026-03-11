package presentation

import "github.com/gin-gonic/gin"

// RegisterOrderRoutes registers order module routes
func RegisterOrderRoutes(r *gin.RouterGroup, handler *OrderHandler) {
	ordersRouter := r.Group("/orders")
	{
		ordersRouter.POST("", handler.CreateOrderFromCart)          // Create order from cart
		ordersRouter.GET("", handler.ListOrders)                   // List user's orders
		ordersRouter.GET("/:order_id", handler.GetOrder)           // Get order details
		ordersRouter.DELETE("/:order_id", handler.CancelOrder)     // Cancel order
	}
}
