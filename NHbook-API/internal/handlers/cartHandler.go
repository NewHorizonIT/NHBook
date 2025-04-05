package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService services.ICartService
	bookService services.IBookService
}

func NewCartHandler(cs services.ICartService, bs services.IBookService) *CartHandler {
	return &CartHandler{
		cartService: cs,
		bookService: bs,
	}
}

func (ch *CartHandler) AddItemToCart(c *gin.Context) {
	// Step 1: Get userID
	userID := c.GetString("userID")

	// Step 2: Binding body
	var item models.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Body invalid")
		return
	}

	// Step 3: Check product is exist

	isExist, err := ch.bookService.CheckBookExist(int(item.ID))

	if err != nil || !isExist {
		utils.WriteError(c, http.StatusBadRequest, "book not found")
		return
	}

	// Step 4: Add to cart
	if err := ch.cartService.AddItemCartByID(userID, &item); err != nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Sprintf("add to cart error %v", err))
		return
	}

	// Step 5: Get cart
	cart, err := ch.cartService.GetCartByID(userID)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Get cart error")
		return
	}

	// Step 6: create response
	utils.WriteResponse(c, http.StatusOK, "Add to cart success", cart, nil)

}

func (ch *CartHandler) GetCart(c *gin.Context) {
	userID := c.GetString("userID")

	cart, err := ch.cartService.GetCartByID(userID)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Get car unsuccess")
		return
	}

	utils.WriteResponse(c, http.StatusOK, "Get cart Success", cart, nil)
}

func (ch *CartHandler) RemoveAllItemToCart(c *gin.Context) {
	userID := c.GetString("userID")

	cart, err := ch.cartService.DeleteCartByID(userID)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Sprintf("Delete all item Cart Error %v", err))
		return
	}

	utils.WriteResponse(c, http.StatusOK, "Delete all item in cart success", cart, nil)
}

func (ch *CartHandler) RemoveItemInCart(c *gin.Context) {
	// Step 1: Get userID and bookID
	userID := c.GetString("userID")

	if userID == "" {
		utils.WriteError(c, http.StatusBadRequest, "Get userID error")
		return
	}

	bookID := c.Param("bookID")

	if bookID == "" {
		utils.WriteError(c, http.StatusBadRequest, "Not Found book ID")
		return
	}

	// Step 2: Convert bookId to integer
	bookIDInt, err := strconv.Atoi(bookID)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("convert book ID error: %v", err).Error())
		return
	}

	// Step 3: Check bookID is exist
	isExist, err := ch.bookService.CheckBookExist(bookIDInt)

	if err != nil || !isExist {
		utils.WriteError(c, http.StatusBadRequest, "Not found Book")
		return
	}

	// Step 4: Remove item in cart
	cart, err := ch.cartService.RemoveItemCart(userID, bookIDInt)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("remove item error: %v", err).Error())
		return
	}

	// Step 5: Return response
	utils.WriteResponse(c, http.StatusOK, "Remove book Success", cart, nil)
}
