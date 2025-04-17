package handlers

import (
	"net/http"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.IOrderService
}

func NewOrderHandler(os services.IOrderService) *OrderHandler {
	return &OrderHandler{
		orderService: os,
	}
}

func (oh *OrderHandler) CreateOrder(c *gin.Context) {
	// Bind the request body to the OrderRequest struct
	var orderRequest request.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call the service to create the order
	orderResponse, err := oh.orderService.CreateOrder(&orderRequest)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteResponse(c, http.StatusOK, "Create Order success", orderResponse, nil)

}
