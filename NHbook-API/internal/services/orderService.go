package services

import (
	"errors"
	"fmt"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
)

var (
	ErrOutOfStock = errors.New("out of stock")
)

type IOrderService interface {
	CreateOrder(order *request.OrderRequest) (*response.OrderResponse, error)
}

type orderService struct {
	orderRepo repositories.IOrderRepository
	userRepo  repositories.IUserRepository
	bookRepo  repositories.IBookRepository
	cartRepo  repositories.ICartRepository
}

// CreateOrder implements IOrderService.
func (o *orderService) CreateOrder(order *request.OrderRequest) (*response.OrderResponse, error) {
	// Step 1: Get cart by userID
	cart, err := o.cartRepo.GetCart(order.UserID)

	if err != nil {
		return nil, err
	}

	// Step 2: Check book stock in warehouse and calculator total Amount
	var totalAmount int = 0
	for i, _ := range cart {
		stock, err := o.bookRepo.GetStock(int(cart[i].ID))
		if err != nil {
			return nil, err
		}
		if stock < cart[i].Quantity {
			return nil, ErrOutOfStock
		}
		totalAmount += cart[i].Total
	}

	// Step 3: Create Order Item

	tx := global.MySQL.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	fmt.Println("UserID", order.UserID)

	newOrder := &models.Order{
		UserID:        order.UserID,
		TotalAmount:   totalAmount,
		Status:        global.OrderStatusPending,
		PaymentMethod: order.PaymentMethod,
	}

	fmt.Printf("Before insert %v\n", newOrder.UserID)

	if err := o.orderRepo.CreateOrder(tx, newOrder); err != nil {
		tx.Rollback()
		return nil, err
	}

	orderItemResponse := make([]response.OrderItemResponse, len(cart))
	// Step 4: Create Order Item
	for i, _ := range cart {
		// Step 4.1: Create Order Item
		orderItem := &models.OrderItem{
			OrderID:  newOrder.ID,
			BookID:   cart[i].ID,
			Quantity: cart[i].Quantity,
			Price:    cart[i].Price,
		}
		if err := o.orderRepo.CreateOrderItem(tx, orderItem); err != nil {
			tx.Rollback()
			return nil, err
		}
		title, err := o.bookRepo.GetTitleBookByID(tx, orderItem.BookID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		orderItemResponse[i] = response.OrderItemResponse{
			BookID:   orderItem.BookID,
			Price:    orderItem.Price,
			BookName: title,
			Quantity: orderItem.Quantity,
			Total:    orderItem.Price * orderItem.Quantity,
		}

		// Step 4.2: Update Book Stock
		if err := o.bookRepo.UpdateStock(tx, int(cart[i].ID), cart[i].Quantity); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Step 5: Create Order Response
	orderResponse := &response.OrderResponse{
		UserID:        newOrder.UserID,
		PaymentMethod: newOrder.PaymentMethod,
		TotalAmount:   newOrder.TotalAmount,
		Status:        newOrder.Status,
		CreatedAt:     newOrder.CreatedAt.Format("2006-01-02 15:04:05"),
		OrderItems:    orderItemResponse,
	}
	return orderResponse, nil

}

func NewOrderService(or repositories.IOrderRepository, ur repositories.IUserRepository, br repositories.IBookRepository, cr repositories.ICartRepository) IOrderService {
	return &orderService{
		orderRepo: or,
		userRepo:  ur,
		bookRepo:  br,
		cartRepo:  cr,
	}
}
