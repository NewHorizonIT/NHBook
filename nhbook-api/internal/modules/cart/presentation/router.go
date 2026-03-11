package presentation

import "github.com/gin-gonic/gin"

// RegisterCartRoutes registers cart module routes
func RegisterCartRoutes(r *gin.RouterGroup, handler *CartHandler) {
	cartRouter := r.Group("/cart")
	{
		cartRouter.GET("", handler.GetCart)                          // Get cart
		cartRouter.POST("/items", handler.AddToCart)                 // Add item to cart
		cartRouter.PUT("/items", handler.UpdateCartItem)             // Update item quantity
		cartRouter.DELETE("/items/:book_id", handler.RemoveFromCart) // Remove item from cart
		cartRouter.DELETE("", handler.ClearCart)                     // Clear cart
	}
}
