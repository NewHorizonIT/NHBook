package domain

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/shared"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ============================================================================
// AUTHOR REPOSITORY
// ============================================================================

type IAuthorRepository interface {
	Create(ctx context.Context, author *Author) error
	GetByID(ctx context.Context, id uuid.UUID) (*Author, error)
	GetAll(ctx context.Context, pagination shared.Pagination) ([]Author, error)
	GetCount(ctx context.Context) (int64, error)
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

// ============================================================================
// CATEGORY REPOSITORY
// ============================================================================

type ICategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	GetAll(ctx context.Context, pagination shared.Pagination) ([]Category, error)
	GetAllWithoutPagination(ctx context.Context) ([]Category, error)
	GetCount(ctx context.Context) (int64, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}

// ============================================================================
// PUBLISHER REPOSITORY
// ============================================================================

type IPublisherRepository interface {
	Create(ctx context.Context, publisher *Publisher) error
	GetByID(ctx context.Context, id uuid.UUID) (*Publisher, error)
	GetAll(ctx context.Context, pagination shared.Pagination) ([]Publisher, error)
	GetAllWithoutPagination(ctx context.Context) ([]Publisher, error)
	GetCount(ctx context.Context) (int64, error)
	Update(ctx context.Context, publisher *Publisher) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}

// ============================================================================
// BOOK REPOSITORY
// ============================================================================

type BookFilter struct {
	CategoryID  *uuid.UUID
	AuthorID    *uuid.UUID
	PublisherID *uuid.UUID
	IsActive    *bool
	MinPrice    *decimal.Decimal
	MaxPrice    *decimal.Decimal
	SearchQuery *string
}

type IBookRepository interface {
	// CRUD
	Create(ctx context.Context, book *Book) error
	GetByID(ctx context.Context, id uuid.UUID) (*Book, error)
	GetByIDWithRelations(ctx context.Context, id uuid.UUID) (*BookWithRelations, error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)

	// List & Filter
	GetAll(ctx context.Context, pagination shared.Pagination) ([]BookWithRelations, error)
	GetCount(ctx context.Context) (int64, error)
	GetActive(ctx context.Context, pagination shared.Pagination) ([]BookWithRelations, error)
	GetActiveCount(ctx context.Context) (int64, error)
	GetByCategory(ctx context.Context, categoryID uuid.UUID, pagination shared.Pagination) ([]BookWithRelations, error)
	GetByAuthor(ctx context.Context, authorID uuid.UUID, pagination shared.Pagination) ([]BookWithRelations, error)
	GetByPublisher(ctx context.Context, publisherID uuid.UUID, pagination shared.Pagination) ([]BookWithRelations, error)
	GetByPriceRange(ctx context.Context, minPrice, maxPrice decimal.Decimal, pagination shared.Pagination) ([]BookWithRelations, error)

	// Search
	Search(ctx context.Context, query string, pagination shared.Pagination) ([]BookWithRelations, error)
	SearchCount(ctx context.Context, query string) (int64, error)

	// Stock management
	AdjustStock(ctx context.Context, id uuid.UUID, delta int32) (*Book, error)
	CheckStock(ctx context.Context, id uuid.UUID) (int32, error)
	GetLowStock(ctx context.Context, threshold int32, pagination shared.Pagination) ([]BookWithRelations, error)

	// Price management
	ChangePrice(ctx context.Context, id uuid.UUID, newPrice decimal.Decimal) (*Book, error)

	// Status management
	ToggleActive(ctx context.Context, id uuid.UUID) (*Book, error)
	SetActive(ctx context.Context, id uuid.UUID, isActive bool) (*Book, error)
}
