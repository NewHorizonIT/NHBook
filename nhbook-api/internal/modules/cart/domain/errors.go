package domain

import "errors"

var (
	ErrCartNotFound    = errors.New("cart not found")
	ErrItemNotFound    = errors.New("item not found in cart")
	ErrInvalidQuantity = errors.New("quantity must be greater than 0")
	ErrCartEmpty       = errors.New("cart is empty")
)
