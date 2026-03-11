package presentation

import (
	"net/http"

	_ "github.com/NewHorizonIT/nhbook-api/internal/modules/order/application"
	usecase "github.com/NewHorizonIT/nhbook-api/internal/modules/order/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/middlewares"
	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/gin-gonic/gin"
	validatorpkg "github.com/go-playground/validator/v10"
)

type OrderHandler struct {
	createOrderUsecase *usecase.CreateOrderFromCartUsecase
	getOrderUsecase    *usecase.GetOrderUsecase
	listOrdersUsecase  *usecase.ListOrdersUsecase
	cancelOrderUsecase *usecase.CancelOrderUsecase
	validator          *validatorpkg.Validate
}

func NewOrderHandler(
	createOrderUsecase *usecase.CreateOrderFromCartUsecase,
	getOrderUsecase *usecase.GetOrderUsecase,
	listOrdersUsecase *usecase.ListOrdersUsecase,
	cancelOrderUsecase *usecase.CancelOrderUsecase,
) *OrderHandler {
	return &OrderHandler{
		createOrderUsecase: createOrderUsecase,
		getOrderUsecase:    getOrderUsecase,
		listOrdersUsecase:  listOrdersUsecase,
		cancelOrderUsecase: cancelOrderUsecase,
		validator:          validatorpkg.New(),
	}
}

// CreateOrderFromCart godoc
// @Summary      Create order from cart
// @Description  Create a new order from user's shopping cart
// @Tags         orders
// @Accept       json
// @Produce      json
// @Success      201 {object} application.CreateOrderResponseDTO
// @Router       /orders [post]
// @Security     BearerAuth
func (h *OrderHandler) CreateOrderFromCart(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		domainErr := errs.NewDomainError(
			errs.CodeUnauthorized,
			"User ID not found in context",
			http.StatusUnauthorized,
		)
		middlewares.HandleError(c, domainErr)
		return
	}

	ctx := c.Request.Context()
	result, err := h.createOrderUsecase.Execute(ctx, userID.(string))
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    result,
	})
}

// GetOrder godoc
// @Summary      Get order details
// @Description  Retrieve order details by order ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order_id path string true "Order ID"
// @Success      200 {object} application.OrderDTO
// @Router       /orders/{order_id} [get]
// @Security     BearerAuth
func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID := c.Param("order_id")

	ctx := c.Request.Context()
	result, err := h.getOrderUsecase.Execute(ctx, orderID)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// ListOrders godoc
// @Summary      List user's orders
// @Description  Retrieve paginated list of user's orders
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        limit query int false "Items per page" default(10)
// @Success      200 {object} application.ListOrdersDTO
// @Router       /orders [get]
// @Security     BearerAuth
func (h *OrderHandler) ListOrders(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		domainErr := errs.NewDomainError(
			errs.CodeUnauthorized,
			"User ID not found in context",
			http.StatusUnauthorized,
		)
		middlewares.HandleError(c, domainErr)
		return
	}

	// Get pagination params
	page := 1
	limit := 10
	c.BindQuery(&struct {
		Page  *int `form:"page" binding:"omitempty,min=1"`
		Limit *int `form:"limit" binding:"omitempty,min=1,max=100"`
	}{&page, &limit})

	ctx := c.Request.Context()
	result, err := h.listOrdersUsecase.Execute(ctx, userID.(string), page, limit)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// CancelOrder godoc
// @Summary      Cancel order
// @Description  Cancel a pending or paid order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order_id path string true "Order ID"
// @Success      204
// @Router       /orders/{order_id} [delete]
// @Security     BearerAuth
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	orderID := c.Param("order_id")

	ctx := c.Request.Context()
	err := h.cancelOrderUsecase.Execute(ctx, orderID)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
