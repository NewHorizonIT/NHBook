package domain

import (
	"context"

	"github.com/google/uuid"
)

type IOrderRepository interface {
	// CreateOrder creates a new order
	CreateOrder(ctx context.Context, order *Order) error

	// GetOrderByID retrieves order by ID
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*Order, error)

	// GetOrdersByUserID retrieves all orders for a user
	GetOrdersByUserID(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*Order, int, error)

	// UpdateOrderStatus updates order status
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status OrderStatus) error

	// AddOrderItem adds item to order
	AddOrderItem(ctx context.Context, orderID uuid.UUID, item *OrderItem) error

	// CancelOrder cancels an order
	CancelOrder(ctx context.Context, orderID uuid.UUID) error
}
