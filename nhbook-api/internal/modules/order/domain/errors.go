package domain

import "errors"

var (
	ErrOrderNotFound        = errors.New("order not found")
	ErrInvalidQuantity      = errors.New("quantity must be greater than 0")
	ErrInvalidPrice         = errors.New("price must be positive")
	ErrCannotCancelShipped  = errors.New("cannot cancel shipped orders")
	ErrCannotModifyPaidOrder = errors.New("cannot modify paid orders")
	ErrOrderEmpty           = errors.New("order is empty")
)
