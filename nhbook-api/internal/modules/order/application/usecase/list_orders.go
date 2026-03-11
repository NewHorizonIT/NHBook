package application

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/order/domain"
	"github.com/google/uuid"
)

type ListOrdersUsecase struct {
	repo domain.IOrderRepository
}

func NewListOrdersUsecase(repo domain.IOrderRepository) *ListOrdersUsecase {
	return &ListOrdersUsecase{
		repo: repo,
	}
}

func (uc *ListOrdersUsecase) Execute(
	ctx context.Context,
	userID string,
	page int,
	limit int,
) (*application.ListOrdersDTO, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Get orders
	orders, total, err := uc.repo.GetOrdersByUserID(ctx, parsedUserID, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs
	orderDTOs := make([]application.OrderDTO, 0)
	for _, order := range orders {
		orderDTOs = append(orderDTOs, *mapOrderToDTO(order))
	}

	totalPages := (total + limit - 1) / limit

	return &application.ListOrdersDTO{
		Orders:     orderDTOs,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}
