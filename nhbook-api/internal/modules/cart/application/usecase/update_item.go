package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/domain"
	"github.com/google/uuid"
)

type UpdateCartItemUsecase struct {
	repo domain.ICartRepository
}

func NewUpdateCartItemUsecase(repo domain.ICartRepository) *UpdateCartItemUsecase {
	return &UpdateCartItemUsecase{
		repo: repo,
	}
}

func (uc *UpdateCartItemUsecase) Execute(ctx context.Context, userID string, dto *application.UpdateCartItemDTO) (*application.CartDTO, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Get cart
	cart, err := uc.repo.GetCart(ctx, parsedUserID)
	if err != nil {
		return nil, err
	}

	// Update in repository
	err = uc.repo.UpdateItemQuantity(ctx, cart.ID, dto.BookID, dto.Quantity)
	if err != nil {
		return nil, err
	}

	// Update in cart state
	err = cart.UpdateItemQuantity(dto.BookID, dto.Quantity)
	if err != nil {
		return nil, err
	}

	return mapCartToDTO(cart), nil
}
