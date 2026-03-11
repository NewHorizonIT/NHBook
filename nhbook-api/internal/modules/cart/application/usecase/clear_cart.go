package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/domain"
	"github.com/google/uuid"
)

type ClearCartUsecase struct {
	repo domain.ICartRepository
}

func NewClearCartUsecase(repo domain.ICartRepository) *ClearCartUsecase {
	return &ClearCartUsecase{
		repo: repo,
	}
}

func (uc *ClearCartUsecase) Execute(ctx context.Context, userID string) error {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	// Get cart
	cart, err := uc.repo.GetCart(ctx, parsedUserID)
	if err != nil {
		return err
	}

	// Clear in repository
	err = uc.repo.ClearCart(ctx, cart.ID)
	if err != nil {
		return err
	}

	// Clear in cart state
	cart.Clear()

	return nil
}
