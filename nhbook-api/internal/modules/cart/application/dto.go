package application

import "github.com/google/uuid"

// Request DTOs
type AddToCartDTO struct {
	BookID   uuid.UUID `json:"book_id" validate:"required"`
	Quantity int       `json:"quantity" validate:"required,gt=0"`
}

type RemoveFromCartDTO struct {
	BookID uuid.UUID `json:"book_id" validate:"required"`
}

type UpdateCartItemDTO struct {
	BookID   uuid.UUID `json:"book_id" validate:"required"`
	Quantity int       `json:"quantity" validate:"required,gt=0"`
}

// Response DTOs
type CartItemDTO struct {
	ID        string `json:"id"`
	BookID    string `json:"book_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CartDTO struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	Items     []CartItemDTO `json:"items"`
	ItemCount int           `json:"item_count"`
	Total     int           `json:"total"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}

type AddToCartResponseDTO struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	Items     []CartItemDTO `json:"items"`
	ItemCount int           `json:"item_count"`
	Total     int           `json:"total"`
	UpdatedAt string        `json:"updated_at"`
}
