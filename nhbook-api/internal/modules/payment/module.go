package payment

import (
	"database/sql"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/infrastructure"
	presentation "github.com/NewHorizonIT/nhbook-api/internal/modules/payment/presentation"
	"github.com/gin-gonic/gin"
)

// PaymentModule manages payment module
type PaymentModule struct {
	handler *presentation.PaymentHandler
}

// NewPaymentModule initializes payment module with all dependencies
func NewPaymentModule(db *sql.DB) *PaymentModule {
	// 1. Infrastructure Layer - Repository
	paymentRepo := infrastructure.NewPaymentRepository(db)

	// 2. Application Layer - Use Cases
	createPaymentUsecase := usecase.NewCreatePaymentUsecase(paymentRepo, nil) // TODO: implement OrderPort
	processPaymentUsecase := usecase.NewProcessPaymentUsecase(paymentRepo, nil) // TODO: implement PaymentGateway
	getPaymentStatusUsecase := usecase.NewGetPaymentStatusUsecase(paymentRepo)

	// 3. Presentation Layer - Handler
	handler := presentation.NewPaymentHandler(
		createPaymentUsecase,
		processPaymentUsecase,
		getPaymentStatusUsecase,
	)

	return &PaymentModule{
		handler: handler,
	}
}

// RegisterRoutes registers routes for payment module
func (m *PaymentModule) RegisterRoutes(router *gin.RouterGroup) {
	presentation.RegisterPaymentRoutes(router, m.handler)
}

// GetHandler returns handler instance
func (m *PaymentModule) GetHandler() *presentation.PaymentHandler {
	return m.handler
}