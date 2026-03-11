package application

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/domain"
	"github.com/google/uuid"
)

type CancelOrderUsecase struct {
	repo domain.IOrderRepository
}

func NewCancelOrderUsecase(repo domain.IOrderRepository) *CancelOrderUsecase {
	return &CancelOrderUsecase{
		repo: repo,
	}
}

func (uc *CancelOrderUsecase) Execute(ctx context.Context, orderID string) error {
	parsedOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return err
	}

	// Get order
	order, err := uc.repo.GetOrderByID(ctx, parsedOrderID)
	if err != nil {
		return err
	}

	// Cancel order
	err = order.Cancel()
	if err != nil {
		return err
	}

	// Update status in repository
	return uc.repo.UpdateOrderStatus(ctx, parsedOrderID, order.Status)
}
