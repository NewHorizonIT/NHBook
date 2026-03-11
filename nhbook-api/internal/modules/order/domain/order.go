package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

type OrderItem struct {
	ID               uuid.UUID
	OrderID          uuid.UUID
	BookID           uuid.UUID
	BookTitleSnapshot string
	PriceSnapshot    decimal.Decimal
	Quantity         int
	CreatedAt        time.Time
}

func NewOrderItem(bookID uuid.UUID, bookTitle string, price decimal.Decimal, quantity int) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	if price.IsNegative() {
		return nil, ErrInvalidPrice
	}

	return &OrderItem{
		ID:                uuid.New(),
		BookID:            bookID,
		BookTitleSnapshot: bookTitle,
		PriceSnapshot:     price,
		Quantity:          quantity,
		CreatedAt:         time.Now(),
	}, nil
}

type Order struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Items       []*OrderItem
	TotalAmount decimal.Decimal
	Status      OrderStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewOrder(userID uuid.UUID) *Order {
	return &Order{
		ID:          uuid.New(),
		UserID:      userID,
		Items:       make([]*OrderItem, 0),
		TotalAmount: decimal.Zero,
		Status:      OrderStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (o *Order) AddItem(item *OrderItem) {
	item.OrderID = o.ID
	o.Items = append(o.Items, item)
	// Calculate total: price * quantity
	itemTotal := item.PriceSnapshot.Mul(decimal.NewFromInt(int64(item.Quantity)))
	o.TotalAmount = o.TotalAmount.Add(itemTotal)
	o.UpdatedAt = time.Now()
}

func (o *Order) Cancel() error {
	if o.Status == OrderStatusShipped {
		return ErrCannotCancelShipped
	}
	o.Status = OrderStatusCancelled
	o.UpdatedAt = time.Now()
	return nil
}

func (o *Order) MarkAsPaid() {
	o.Status = OrderStatusPaid
	o.UpdatedAt = time.Now()
}

func (o *Order) MarkAsShipped() {
	o.Status = OrderStatusShipped
	o.UpdatedAt = time.Now()
}

func (o *Order) GetItemCount() int {
	return len(o.Items)
}

func (o *Order) IsEmpty() bool {
	return len(o.Items) == 0
}

func (o *Order) IsPaid() bool {
	return o.Status == OrderStatusPaid || o.Status == OrderStatusShipped
}
