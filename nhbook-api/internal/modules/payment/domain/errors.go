package domain

import "errors"

var (
	ErrPaymentNotFound         = errors.New("payment not found")
	ErrDuplicateIdempotencyKey = errors.New("idempotency key already used")
	ErrInvalidAmount           = errors.New("amount must be positive")
	ErrInvalidIdempotencyKey   = errors.New("idempotency key cannot be empty")
	ErrCannotModifyPayment     = errors.New("cannot modify completed payment")
)
