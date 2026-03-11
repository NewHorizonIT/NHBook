package domain

import (
	"context"

	"github.com/google/uuid"
)

type ICartRepository interface {
	// GetCart retrieves user's cart
	GetCart(ctx context.Context, userID uuid.UUID) (*Cart, error)

	// CreateCart creates a new cart for user
	CreateCart(ctx context.Context, cart *Cart) error

	// AddItemToCart adds or updates item in cart
	AddItemToCart(ctx context.Context, cartID uuid.UUID, item *CartItem) error

	// RemoveItemFromCart removes item from cart
	RemoveItemFromCart(ctx context.Context, cartID uuid.UUID, bookID uuid.UUID) error

	// UpdateItemQuantity updates quantity of item in cart
	UpdateItemQuantity(ctx context.Context, cartID uuid.UUID, bookID uuid.UUID, quantity int) error

	// ClearCart removes all items from cart
	ClearCart(ctx context.Context, cartID uuid.UUID) error

	// DeleteCart deletes entire cart
	DeleteCart(ctx context.Context, cartID uuid.UUID) error
}
