package domain

import "errors"

var (
	// Book errors
	ErrBookNotFound        = errors.New("book not found")
	ErrBookTitleRequired   = errors.New("book title is required")
	ErrBookPriceInvalid    = errors.New("book price must be greater than 0")
	ErrBookStockInvalid    = errors.New("book stock cannot be negative")
	ErrInsufficientStock   = errors.New("insufficient stock")
	ErrBookAlreadyInactive = errors.New("book is already inactive")
	ErrBookAlreadyActive   = errors.New("book is already active")

	// Author errors
	ErrAuthorNotFound     = errors.New("author not found")
	ErrAuthorNameRequired = errors.New("author name is required")
	ErrAuthorHasBooks     = errors.New("cannot delete author with existing books")

	// Category errors
	ErrCategoryNotFound     = errors.New("category not found")
	ErrCategoryNameRequired = errors.New("category name is required")
	ErrCategoryNameExists   = errors.New("category name already exists")
	ErrCategoryHasBooks     = errors.New("cannot delete category with existing books")

	// Publisher errors
	ErrPublisherNotFound     = errors.New("publisher not found")
	ErrPublisherNameRequired = errors.New("publisher name is required")
	ErrPublisherHasBooks     = errors.New("cannot delete publisher with existing books")
)
