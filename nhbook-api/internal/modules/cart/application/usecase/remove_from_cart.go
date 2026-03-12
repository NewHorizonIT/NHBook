package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/domain"
	"github.com/google/uuid"
)

type RemoveFromCartUsecase struct {
	repo domain.ICartRepository
}

func NewRemoveFromCartUsecase(repo domain.ICartRepository) *RemoveFromCartUsecase {
	return &RemoveFromCartUsecase{
		repo: repo,
	}
}

func (uc *RemoveFromCartUsecase) Execute(ctx context.Context, userID string, bookID uuid.UUID) error {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	// Get cart
	cart, err := uc.repo.GetCart(ctx, parsedUserID)
	if err != nil {
		return err
	}

	// Remove from repository
	err = uc.repo.RemoveItemFromCart(ctx, cart.ID, bookID)
	if err != nil {
		return err
	}

	// Remove from cart state
	_ = cart.RemoveItem(bookID)

	return nil
}
