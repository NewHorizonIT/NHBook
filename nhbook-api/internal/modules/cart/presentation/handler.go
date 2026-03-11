package presentation

import (
	"net/http"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/application/usecase"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/middlewares"
	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/gin-gonic/gin"
	validatorpkg "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CartHandler struct {
	getCartUsecase        *usecase.GetCartUsecase
	addToCartUsecase      *usecase.AddToCartUsecase
	removeFromCartUsecase *usecase.RemoveFromCartUsecase
	updateCartItemUsecase *usecase.UpdateCartItemUsecase
	clearCartUsecase      *usecase.ClearCartUsecase
	validator             *validatorpkg.Validate
}

func NewCartHandler(
	getCartUsecase *usecase.GetCartUsecase,
	addToCartUsecase *usecase.AddToCartUsecase,
	removeFromCartUsecase *usecase.RemoveFromCartUsecase,
	updateCartItemUsecase *usecase.UpdateCartItemUsecase,
	clearCartUsecase *usecase.ClearCartUsecase,
) *CartHandler {
	return &CartHandler{
		getCartUsecase:        getCartUsecase,
		addToCartUsecase:      addToCartUsecase,
		removeFromCartUsecase: removeFromCartUsecase,
		updateCartItemUsecase: updateCartItemUsecase,
		clearCartUsecase:      clearCartUsecase,
		validator:             validatorpkg.New(),
	}
}

// GetCart godoc
// @Summary      Get user's cart
// @Description  Retrieve the contents of user's shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Success      200 {object} application.CartDTO
// @Router       /cart [get]
// @Security     BearerAuth
func (h *CartHandler) GetCart(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
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
	result, err := h.getCartUsecase.Execute(ctx, userID.(string))
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// AddToCart godoc
// @Summary      Add item to cart
// @Description  Add a book to user's shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        request body application.AddToCartDTO true "Item to add"
// @Success      201 {object} application.AddToCartResponseDTO
// @Router       /cart/items [post]
// @Security     BearerAuth
func (h *CartHandler) AddToCart(c *gin.Context) {
	var req application.AddToCartDTO

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

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		).WithError(err)
		middlewares.HandleError(c, domainErr)
		return
	}

	// Validate
	if err := h.validator.Struct(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeValidationError,
			"Validation failed",
			http.StatusBadRequest,
		)
		middlewares.HandleError(c, domainErr)
		return
	}

	ctx := c.Request.Context()
	result, err := h.addToCartUsecase.Execute(ctx, userID.(string), &req)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    result,
	})
}

// RemoveFromCart godoc
// @Summary      Remove item from cart
// @Description  Remove a book from user's shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        book_id path string true "Book ID"
// @Success      204
// @Router       /cart/items/{book_id} [delete]
// @Security     BearerAuth
func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	bookID := c.Param("book_id")

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

	// Parse book ID
	parsedBookID, err := uuid.Parse(bookID)
	if err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeInvalidInput,
			"Invalid book ID format",
			http.StatusBadRequest,
		).WithError(err)
		middlewares.HandleError(c, domainErr)
		return
	}

	ctx := c.Request.Context()
	err = h.removeFromCartUsecase.Execute(ctx, userID.(string), parsedBookID)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// UpdateCartItem godoc
// @Summary      Update item quantity in cart
// @Description  Update the quantity of a book in user's shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        request body application.UpdateCartItemDTO true "Item quantity to update"
// @Success      200 {object} application.CartDTO
// @Router       /cart/items [put]
// @Security     BearerAuth
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	var req application.UpdateCartItemDTO

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

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeInvalidInput,
			"Invalid request body",
			http.StatusBadRequest,
		).WithError(err)
		middlewares.HandleError(c, domainErr)
		return
	}

	// Validate
	if err := h.validator.Struct(&req); err != nil {
		domainErr := errs.NewDomainError(
			errs.CodeValidationError,
			"Validation failed",
			http.StatusBadRequest,
		)
		middlewares.HandleError(c, domainErr)
		return
	}

	ctx := c.Request.Context()
	result, err := h.updateCartItemUsecase.Execute(ctx, userID.(string), &req)
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// ClearCart godoc
// @Summary      Clear cart
// @Description  Remove all items from user's shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Success      204
// @Router       /cart [delete]
// @Security     BearerAuth
func (h *CartHandler) ClearCart(c *gin.Context) {
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
	err := h.clearCartUsecase.Execute(ctx, userID.(string))
	if err != nil {
		middlewares.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
