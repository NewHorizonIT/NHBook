package domain

import (
	"context"

	"github.com/google/uuid"
)

type IPaymentRepository interface {
	// CreatePayment creates a new payment
	CreatePayment(ctx context.Context, payment *Payment) error

	// GetPaymentByID retrieves payment by ID
	GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (*Payment, error)

	// GetPaymentByOrderID retrieves payment by order ID
	GetPaymentByOrderID(ctx context.Context, orderID uuid.UUID) (*Payment, error)

	// GetPaymentByIdempotencyKey retrieves payment by idempotency key
	GetPaymentByIdempotencyKey(ctx context.Context, key string) (*Payment, error)

	// UpdatePaymentStatus updates payment status
	UpdatePaymentStatus(ctx context.Context, paymentID uuid.UUID, status PaymentStatus) error
}
