package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/domain"
	"github.com/google/uuid"
)

type AddToCartUsecase struct {
	repo domain.ICartRepository
}

func NewAddToCartUsecase(repo domain.ICartRepository) *AddToCartUsecase {
	return &AddToCartUsecase{
		repo: repo,
	}
}

func (uc *AddToCartUsecase) Execute(ctx context.Context, userID string, dto *application.AddToCartDTO) (*application.AddToCartResponseDTO, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Get or create cart
	cart, err := uc.repo.GetCart(ctx, parsedUserID)
	if err != nil {
		// If cart doesn't exist, create new one
		if err == domain.ErrCartNotFound {
			cart = domain.NewCart(parsedUserID)
			if err := uc.repo.CreateCart(ctx, cart); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Create new cart item
	newItem, err := domain.NewCartItem(cart.ID, dto.BookID, dto.Quantity)
	if err != nil {
		return nil, err
	}

	// Add to repository
	err = uc.repo.AddItemToCart(ctx, cart.ID, newItem)
	if err != nil {
		return nil, err
	}

	// Add to cart state
	_ = cart.AddItem(newItem)

	// Convert to response DTO
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

	return &application.AddToCartResponseDTO{
		ID:        cart.ID.String(),
		UserID:    cart.UserID.String(),
		Items:     items,
		ItemCount: cart.GetItemCount(),
		Total:     cart.GetTotalQuantity(),
		UpdatedAt: cart.UpdatedAt.String(),
	}, nil
}
