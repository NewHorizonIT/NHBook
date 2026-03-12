package application

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/domain"
	"github.com/google/uuid"
)

type GetOrderUsecase struct {
	repo domain.IOrderRepository
}

func NewGetOrderUsecase(repo domain.IOrderRepository) *GetOrderUsecase {
	return &GetOrderUsecase{
		repo: repo,
	}
}

func (uc *GetOrderUsecase) Execute(ctx context.Context, orderID string) (*application.OrderDTO, error) {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return nil, err
	}

	order, err := uc.repo.GetOrderByID(ctx, parsedOrderID)
	if err != nil {
		return nil, err
	}

	return mapOrderToDTO(order), nil
}

func mapOrderToDTO(order *domain.Order) *application.OrderDTO {
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

	return &application.OrderDTO{
		ID:          order.ID.String(),
		UserID:      order.UserID.String(),
		Items:       items,
		TotalAmount: order.TotalAmount.String(),
		Status:      string(order.Status),
		ItemCount:   order.GetItemCount(),
		CreatedAt:   order.CreatedAt.String(),
		UpdatedAt:   order.UpdatedAt.String(),
	}
}
