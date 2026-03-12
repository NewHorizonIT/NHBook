package application

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CartPort defines the interface to fetch cart from cart module
type CartPort interface {
	GetCartItemsForOrder(ctx context.Context, userID uuid.UUID) ([]CartItemData, error)
}

type CartItemData struct {
	BookID    uuid.UUID
	BookTitle string
	Price     string // Using string for decimal
	Quantity  int
}

type CreateOrderFromCartUsecase struct {
	orderRepo domain.IOrderRepository
	cartPort  CartPort
}

func NewCreateOrderFromCartUsecase(
	orderRepo domain.IOrderRepository,
	cartPort CartPort,
) *CreateOrderFromCartUsecase {
	return &CreateOrderFromCartUsecase{
		orderRepo: orderRepo,
		cartPort:  cartPort,
	}
}

func (uc *CreateOrderFromCartUsecase) Execute(
	ctx context.Context,
	userID string,
) (*application.CreateOrderResponseDTO, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// 1. Get cart items from cart module
	cartItems, err := uc.cartPort.GetCartItemsForOrder(ctx, parsedUserID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, domain.ErrOrderEmpty
	}

	// 2. Create order entity
	order := domain.NewOrder(parsedUserID)

	// 3. Add items to order
	for _, cartItem := range cartItems {
		// Convert price string to decimal
		price, err := parseDecimal(cartItem.Price)
		if err != nil {
			return nil, err
		}

		orderItem, err := domain.NewOrderItem(
			cartItem.BookID,
			cartItem.BookTitle,
			price,
			cartItem.Quantity,
		)
		if err != nil {
			return nil, err
		}

		order.AddItem(orderItem)
	}

	// 4. Save order to repository
	err = uc.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// 5. Convert to response DTO
	items := make([]application.OrderItemDTO, 0)
	for _, item := range order.Items {
		items = append(items, application.OrderItemDTO{
			ID:                item.ID.String(),
			BookID:            item.BookID.String(),
			BookTitleSnapshot: item.BookTitleSnapshot,
			PriceSnapshot:     item.PriceSnapshot.String(),
			Quantity:          item.Quantity,
			CreatedAt:         item.CreatedAt.String(),
		})
	}

	return &application.CreateOrderResponseDTO{
		ID:          order.ID.String(),
		UserID:      order.UserID.String(),
		Items:       items,
		TotalAmount: order.TotalAmount.String(),
		Status:      string(order.Status),
		ItemCount:   order.GetItemCount(),
		CreatedAt:   order.CreatedAt.String(),
	}, nil
}

// Helper function to parse decimal from string
func parseDecimal(s string) (decimal.Decimal, error) {
	return decimal.NewFromString(s)
}
