package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Author struct {
	ID        uuid.UUID
	Name      string
	Bio       *string
	CreatedAt time.Time
}

func NewAuthor(name string, bio *string) (*Author, error) {
	if name == "" {
		return nil, ErrAuthorNameRequired
	}

	return &Author{
		ID:        uuid.New(),
		Name:      name,
		Bio:       bio,
		CreatedAt: time.Now(),
	}, nil
}

func (a *Author) Update(name *string, bio *string) {
	if name != nil {
		a.Name = *name
	}
	if bio != nil {
		a.Bio = bio
	}
}

type Category struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	if name == "" {
		return nil, ErrCategoryNameRequired
	}

	return &Category{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func (c *Category) Update(name string) error {
	if name == "" {
		return ErrCategoryNameRequired
	}
	c.Name = name
	return nil
}

type Publisher struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

func NewPublisher(name string) (*Publisher, error) {
	if name == "" {
		return nil, ErrPublisherNameRequired
	}

	return &Publisher{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func (p *Publisher) Update(name string) error {
	if name == "" {
		return ErrPublisherNameRequired
	}
	p.Name = name
	return nil
}

type Book struct {
	ID          uuid.UUID
	Title       string
	Description *string
	Price       Price
	Stock       Stock
	AuthorID    *uuid.UUID
	CategoryID  *uuid.UUID
	PublisherID *uuid.UUID
	CoverImage  *string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewBookParams struct {
	Title       string
	Description *string
	Price       decimal.Decimal
	Stock       int32
	AuthorID    *uuid.UUID
	CategoryID  *uuid.UUID
	PublisherID *uuid.UUID
	CoverImage  *string
}

func NewBook(params NewBookParams) (*Book, error) {
	if params.Title == "" {
		return nil, ErrBookTitleRequired
	}

	price, err := NewPrice(params.Price)
	if err != nil {
		return nil, err
	}

	stock, err := NewStock(params.Stock)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Book{
		ID:          uuid.New(),
		Title:       params.Title,
		Description: params.Description,
		Price:       price,
		Stock:       stock,
		AuthorID:    params.AuthorID,
		CategoryID:  params.CategoryID,
		PublisherID: params.PublisherID,
		CoverImage:  params.CoverImage,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (b *Book) Update(params UpdateBookParams) error {
	if params.Title != nil {
		if *params.Title == "" {
			return ErrBookTitleRequired
		}
		b.Title = *params.Title
	}

	if params.Description != nil {
		b.Description = params.Description
	}

	if params.Price != nil {
		price, err := NewPrice(*params.Price)
		if err != nil {
			return err
		}
		b.Price = price
	}

	if params.Stock != nil {
		stock, err := NewStock(*params.Stock)
		if err != nil {
			return err
		}
		b.Stock = stock
	}

	if params.AuthorID != nil {
		b.AuthorID = params.AuthorID
	}

	if params.CategoryID != nil {
		b.CategoryID = params.CategoryID
	}

	if params.PublisherID != nil {
		b.PublisherID = params.PublisherID
	}

	if params.CoverImage != nil {
		b.CoverImage = params.CoverImage
	}

	b.UpdatedAt = time.Now()
	return nil
}

type UpdateBookParams struct {
	Title       *string
	Description *string
	Price       *decimal.Decimal
	Stock       *int32
	AuthorID    *uuid.UUID
	CategoryID  *uuid.UUID
	PublisherID *uuid.UUID
	CoverImage  *string
}

func (b *Book) ChangePrice(newPrice decimal.Decimal) error {
	price, err := NewPrice(newPrice)
	if err != nil {
		return err
	}
	b.Price = price
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) AdjustStock(delta int32) error {
	newQty := b.Stock.Value() + delta
	stock, err := NewStock(newQty)
	if err != nil {
		return err
	}
	b.Stock = stock
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) DeductStock(quantity int32) error {
	stock, err := b.Stock.Deduct(quantity)
	if err != nil {
		return err
	}
	b.Stock = stock
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) Activate() error {
	if b.IsActive {
		return ErrBookAlreadyActive
	}
	b.IsActive = true
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) Deactivate() error {
	if !b.IsActive {
		return ErrBookAlreadyInactive
	}
	b.IsActive = false
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) ToggleActive() {
	b.IsActive = !b.IsActive
	b.UpdatedAt = time.Now()
}

func (b *Book) IsInStock() bool {
	return b.Stock.Value() > 0
}

func (b *Book) CanPurchase(quantity int32) bool {
	return b.IsActive && b.Stock.CanDeduct(quantity)
}

type BookWithRelations struct {
	Book
	AuthorName    *string
	CategoryName  *string
	PublisherName *string
}
