package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentStatus string

const (
	PaymentStatusInit    PaymentStatus = "INIT"
	PaymentStatusSuccess PaymentStatus = "SUCCESS"
	PaymentStatusFailed  PaymentStatus = "FAILED"
)

type Payment struct {
	ID             uuid.UUID
	OrderID        uuid.UUID
	Amount         decimal.Decimal
	Status         PaymentStatus
	IdempotencyKey string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewPayment(orderID uuid.UUID, amount decimal.Decimal, idempotencyKey string) (*Payment, error) {
	if amount.IsNegative() || amount.IsZero() {
		return nil, ErrInvalidAmount
	}

	if idempotencyKey == "" {
		return nil, ErrInvalidIdempotencyKey
	}

	return &Payment{
		ID:             uuid.New(),
		OrderID:        orderID,
		Amount:         amount,
		Status:         PaymentStatusInit,
		IdempotencyKey: idempotencyKey,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (p *Payment) MarkAsSuccess() {
	p.Status = PaymentStatusSuccess
	p.UpdatedAt = time.Now()
}

func (p *Payment) MarkAsFailed() {
	p.Status = PaymentStatusFailed
	p.UpdatedAt = time.Now()
}

func (p *Payment) IsSuccessful() bool {
	return p.Status == PaymentStatusSuccess
}

func (p *Payment) IsFailed() bool {
	return p.Status == PaymentStatusFailed
}

func (p *Payment) IsPending() bool {
	return p.Status == PaymentStatusInit
}
