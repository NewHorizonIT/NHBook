package infrastructure

import (
	"context"
	"database/sql"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/cart/infrastructure/sqlc"
	"github.com/google/uuid"
)

type CartRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func NewCartRepository(db *sql.DB) domain.ICartRepository {
	return &CartRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

// GetCart retrieves user's cart with all items
func (r *CartRepository) GetCart(ctx context.Context, userID uuid.UUID) (*domain.Cart, error) {
	// Get cart
	dbCart, err := r.q.GetCartByUserID(ctx, uuid.NullUUID{UUID: userID, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrCartNotFound
		}
		return nil, err
	}

	// Get cart items
	dbItems, err := r.q.GetCartItems(ctx, uuid.NullUUID{UUID: dbCart.ID, Valid: true})
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Convert to domain entities
	cart := &domain.Cart{
		ID:        dbCart.ID,
		UserID:    dbCart.UserID.UUID,
		Items:     make([]*domain.CartItem, 0),
		CreatedAt: dbCart.CreatedAt.Time,
		UpdatedAt: dbCart.UpdatedAt.Time,
	}

	for _, dbItem := range dbItems {
		item := &domain.CartItem{
			ID:        dbItem.ID,
			CartID:    dbItem.CartID.UUID,
			BookID:    dbItem.BookID.UUID,
			Quantity:  int(dbItem.Quantity),
			CreatedAt: dbItem.CreatedAt.Time,
			UpdatedAt: dbItem.UpdatedAt.Time,
		}
		cart.Items = append(cart.Items, item)
	}

	return cart, nil
}

// CreateCart creates a new cart for user
func (r *CartRepository) CreateCart(ctx context.Context, cart *domain.Cart) error {
	return r.q.CreateCart(ctx, sqlc.CreateCartParams{
		ID:        cart.ID,
		UserID:    uuid.NullUUID{UUID: cart.UserID, Valid: true},
		CreatedAt: sql.NullTime{Time: cart.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: cart.UpdatedAt, Valid: true},
	})
}

// AddItemToCart adds or updates item in cart
func (r *CartRepository) AddItemToCart(ctx context.Context, cartID uuid.UUID, item *domain.CartItem) error {
	return r.q.AddCartItem(ctx, sqlc.AddCartItemParams{
		ID:        item.ID,
		CartID:    uuid.NullUUID{UUID: cartID, Valid: true},
		BookID:    uuid.NullUUID{UUID: item.BookID, Valid: true},
		Quantity:  int32(item.Quantity),
		CreatedAt: sql.NullTime{Time: item.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: item.UpdatedAt, Valid: true},
	})
}

// RemoveItemFromCart removes item from cart
func (r *CartRepository) RemoveItemFromCart(ctx context.Context, cartID uuid.UUID, bookID uuid.UUID) error {
	return r.q.RemoveCartItem(ctx, sqlc.RemoveCartItemParams{
		CartID: uuid.NullUUID{UUID: cartID, Valid: true},
		BookID: uuid.NullUUID{UUID: bookID, Valid: true},
	})
}

// UpdateItemQuantity updates quantity of item in cart
func (r *CartRepository) UpdateItemQuantity(ctx context.Context, cartID uuid.UUID, bookID uuid.UUID, quantity int) error {
	return r.q.UpdateCartItemQuantity(ctx, sqlc.UpdateCartItemQuantityParams{
		Quantity:  int32(quantity),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		CartID:    uuid.NullUUID{UUID: cartID, Valid: true},
		BookID:    uuid.NullUUID{UUID: bookID, Valid: true},
	})
}

// ClearCart removes all items from cart
func (r *CartRepository) ClearCart(ctx context.Context, cartID uuid.UUID) error {
	return r.q.ClearCartItems(ctx, uuid.NullUUID{UUID: cartID, Valid: true})
}

// DeleteCart deletes entire cart
func (r *CartRepository) DeleteCart(ctx context.Context, cartID uuid.UUID) error {
	return r.q.DeleteCart(ctx, cartID)
}
