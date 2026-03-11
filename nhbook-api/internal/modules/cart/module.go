package cart

import (
	"database/sql"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/infrastructure"
	presentation "github.com/NewHorizonIT/nhbook-api/internal/modules/cart/presentation"
	"github.com/gin-gonic/gin"
)

// CartModule manages cart module
type CartModule struct {
	handler *presentation.CartHandler
}

// NewCartModule initializes cart module with all dependencies
func NewCartModule(db *sql.DB) *CartModule {
	// 1. Infrastructure Layer - Repository
	cartRepo := infrastructure.NewCartRepository(db)

	// 2. Application Layer - Use Cases
	getCartUsecase := usecase.NewGetCartUsecase(cartRepo)
	addToCartUsecase := usecase.NewAddToCartUsecase(cartRepo)
	removeFromCartUsecase := usecase.NewRemoveFromCartUsecase(cartRepo)
	updateCartItemUsecase := usecase.NewUpdateCartItemUsecase(cartRepo)
	clearCartUsecase := usecase.NewClearCartUsecase(cartRepo)

	// 3. Presentation Layer - Handler
	handler := presentation.NewCartHandler(
		getCartUsecase,
		addToCartUsecase,
		removeFromCartUsecase,
		updateCartItemUsecase,
		clearCartUsecase,
	)

	return &CartModule{
		handler: handler,
	}
}

// RegisterRoutes registers routes for cart module
func (m *CartModule) RegisterRoutes(router *gin.RouterGroup) {
	presentation.RegisterCartRoutes(router, m.handler)
}

// GetHandler returns handler instance (if needed elsewhere)
func (m *CartModule) GetHandler() *presentation.CartHandler {
	return m.handler
}