package presentation

import "github.com/gin-gonic/gin"

// RegisterPaymentRoutes registers payment module routes
func RegisterPaymentRoutes(r *gin.RouterGroup, handler *PaymentHandler) {
	paymentsRouter := r.Group("/payments")
	{
		paymentsRouter.POST("", handler.CreatePayment)                // Create payment
		paymentsRouter.POST("/process", handler.ProcessPayment)       // Process payment
		paymentsRouter.GET("/:payment_id", handler.GetPaymentStatus)  // Get payment status
	}
}
