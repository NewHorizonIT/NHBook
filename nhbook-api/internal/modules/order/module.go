package order

import (
	"database/sql"

	usecase "github.com/NewHorizonIT/nhbook-api/internal/modules/order/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/infrastructure"
	presentation "github.com/NewHorizonIT/nhbook-api/internal/modules/order/presentation"
	"github.com/gin-gonic/gin"
)

// OrderModule manages order module
type OrderModule struct {
	handler *presentation.OrderHandler
}

// NewOrderModule initializes order module with all dependencies
func NewOrderModule(db *sql.DB) *OrderModule {
	// 1. Infrastructure Layer - Repository
	orderRepo := infrastructure.NewOrderRepository(db)

	// 2. Application Layer - Use Cases
	createOrderUsecase := usecase.NewCreateOrderFromCartUsecase(orderRepo, nil) // TODO: implement CartPort
	getOrderUsecase := usecase.NewGetOrderUsecase(orderRepo)
	listOrdersUsecase := usecase.NewListOrdersUsecase(orderRepo)
	cancelOrderUsecase := usecase.NewCancelOrderUsecase(orderRepo)

	// 3. Presentation Layer - Handler
	handler := presentation.NewOrderHandler(
		createOrderUsecase,
		getOrderUsecase,
		listOrdersUsecase,
		cancelOrderUsecase,
	)

	return &OrderModule{
		handler: handler,
	}
}

// RegisterRoutes registers routes for order module
func (m *OrderModule) RegisterRoutes(router *gin.RouterGroup) {
	presentation.RegisterOrderRoutes(router, m.handler)
}

// GetHandler returns handler instance
func (m *OrderModule) GetHandler() *presentation.OrderHandler {
	return m.handler
}
