package application

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ============================================================================
// AUTHOR DTOs
// ============================================================================

type CreateAuthorDTO struct {
	Name string  `json:"name" validate:"required,min=1,max=255"`
	Bio  *string `json:"bio" validate:"omitempty,max=2000"`
}

type UpdateAuthorDTO struct {
	Name *string `json:"name" validate:"omitempty,min=1,max=255"`
	Bio  *string `json:"bio" validate:"omitempty,max=2000"`
}

type AuthorResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Bio       *string   `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type GetAuthorsDTO struct {
	Page     int32 `json:"page" validate:"gte=1"`
	PageSize int32 `json:"page_size" validate:"gte=1,lte=100"`
}

// ============================================================================
// CATEGORY DTOs
// ============================================================================

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type UpdateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type CategoryResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetCategoriesDTO struct {
	Page     int32 `json:"page" validate:"gte=1"`
	PageSize int32 `json:"page_size" validate:"gte=1,lte=100"`
}

// ============================================================================
// PUBLISHER DTOs
// ============================================================================

type CreatePublisherDTO struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type UpdatePublisherDTO struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type PublisherResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPublishersDTO struct {
	Page     int32 `json:"page" validate:"gte=1"`
	PageSize int32 `json:"page_size" validate:"gte=1,lte=100"`
}

// ============================================================================
// BOOK DTOs
// ============================================================================

type CreateBookDTO struct {
	Title       string          `json:"title" validate:"required,min=1,max=255"`
	Description *string         `json:"description" validate:"omitempty,max=5000"`
	Price       decimal.Decimal `json:"price" validate:"required,gt=0"`
	Stock       int32           `json:"stock" validate:"gte=0"`
	AuthorID    *uuid.UUID      `json:"author_id" validate:"omitempty,uuid"`
	CategoryID  *uuid.UUID      `json:"category_id" validate:"omitempty,uuid"`
	PublisherID *uuid.UUID      `json:"publisher_id" validate:"omitempty,uuid"`
	CoverImage  *string         `json:"cover_image" validate:"omitempty,url"`
}

type UpdateBookDTO struct {
	Title       *string          `json:"title" validate:"omitempty,min=1,max=255"`
	Description *string          `json:"description" validate:"omitempty,max=5000"`
	Price       *decimal.Decimal `json:"price" validate:"omitempty,gt=0"`
	Stock       *int32           `json:"stock" validate:"omitempty,gte=0"`
	AuthorID    *uuid.UUID       `json:"author_id" validate:"omitempty,uuid"`
	CategoryID  *uuid.UUID       `json:"category_id" validate:"omitempty,uuid"`
	PublisherID *uuid.UUID       `json:"publisher_id" validate:"omitempty,uuid"`
	CoverImage  *string          `json:"cover_image" validate:"omitempty,url"`
}

type BookResponseDTO struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   *string   `json:"description,omitempty"`
	Price         string    `json:"price"`
	Stock         int32     `json:"stock"`
	AuthorID      *string   `json:"author_id,omitempty"`
	AuthorName    *string   `json:"author_name,omitempty"`
	CategoryID    *string   `json:"category_id,omitempty"`
	CategoryName  *string   `json:"category_name,omitempty"`
	PublisherID   *string   `json:"publisher_id,omitempty"`
	PublisherName *string   `json:"publisher_name,omitempty"`
	CoverImage    *string   `json:"cover_image,omitempty"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AdjustStockDTO struct {
	Delta int32 `json:"delta" validate:"required"`
}

type ChangePriceDTO struct {
	Price decimal.Decimal `json:"price" validate:"required,gt=0"`
}

type SearchBooksDTO struct {
	Query    string `json:"query" validate:"required,min=1"`
	Page     int32  `json:"page" validate:"gte=1"`
	PageSize int32  `json:"page_size" validate:"gte=1,lte=100"`
}

type GetBooksDTO struct {
	Page       int32      `json:"page" validate:"gte=1"`
	PageSize   int32      `json:"page_size" validate:"gte=1,lte=100"`
	CategoryID *uuid.UUID `json:"category_id" validate:"omitempty,uuid"`
	AuthorID   *uuid.UUID `json:"author_id" validate:"omitempty,uuid"`
	ActiveOnly bool       `json:"active_only"`
}

// ============================================================================
// PAGINATION RESPONSE
// ============================================================================

type PaginatedResponse[T any] struct {
	Items      []T   `json:"items"`
	TotalCount int64 `json:"total_count"`
	Page       int32 `json:"page"`
	PageSize   int32 `json:"page_size"`
	TotalPages int32 `json:"total_pages"`
}
