package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/domain"
	"github.com/google/uuid"
)

type GetCartUsecase struct {
	repo domain.ICartRepository
}

func NewGetCartUsecase(repo domain.ICartRepository) *GetCartUsecase {
	return &GetCartUsecase{
		repo: repo,
	}
}

func (uc *GetCartUsecase) Execute(ctx context.Context, userID string) (*application.CartDTO, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	cart, err := uc.repo.GetCart(ctx, parsedUserID)
	if err != nil {
		return nil, err
	}

	// Convert to DTO
	return mapCartToDTO(cart), nil
}

func mapCartToDTO(cart *domain.Cart) *application.CartDTO {
	items := make([]application.CartItemDTO, 0)
	for _, item := range cart.Items {
		items = append(items, application.CartItemDTO{
			ID:        item.ID.String(),
			BookID:    item.BookID.String(),
			Quantity:  item.Quantity,
			CreatedAt: item.CreatedAt.String(),
			UpdatedAt: item.UpdatedAt.String(),
		})
	}

	return &application.CartDTO{
		ID:        cart.ID.String(),
		UserID:    cart.UserID.String(),
		Items:     items,
		ItemCount: cart.GetItemCount(),
		Total:     cart.GetTotalQuantity(),
		CreatedAt: cart.CreatedAt.String(),
		UpdatedAt: cart.UpdatedAt.String(),
	}
}
