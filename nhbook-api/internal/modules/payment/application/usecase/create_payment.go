package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/domain"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderPort defines interface to fetch order from order module
type OrderPort interface {
	GetOrderAmount(ctx context.Context, orderID uuid.UUID) (decimal.Decimal, error)
}

type CreatePaymentUsecase struct {
	paymentRepo domain.IPaymentRepository
	orderPort   OrderPort
}

func NewCreatePaymentUsecase(
	paymentRepo domain.IPaymentRepository,
	orderPort OrderPort,
) *CreatePaymentUsecase {
	return &CreatePaymentUsecase{
		paymentRepo: paymentRepo,
		orderPort:   orderPort,
	}
}

func (uc *CreatePaymentUsecase) Execute(
	ctx context.Context,
	dto *application.CreatePaymentDTO,
) (*application.CreatePaymentResponseDTO, error) {
	// Check idempotency
	existingPayment, err := uc.paymentRepo.GetPaymentByIdempotencyKey(ctx, dto.IdempotencyKey)
	if err == nil && existingPayment != nil {
		// Payment already created with this key, return it
		return &application.CreatePaymentResponseDTO{
			ID:             existingPayment.ID.String(),
			OrderID:        existingPayment.OrderID.String(),
			Amount:         existingPayment.Amount.String(),
			Status:         string(existingPayment.Status),
			IdempotencyKey: existingPayment.IdempotencyKey,
			CreatedAt:      existingPayment.CreatedAt.String(),
		}, nil
	}

	// Parse order ID
	orderID, err := uuid.Parse(dto.OrderID)
	if err != nil {
		return nil, err
	}

	// Get order amount from order service
	amount, err := uc.orderPort.GetOrderAmount(ctx, orderID)
	if err != nil {
		return nil, err
	}

	// Create payment
	payment, err := domain.NewPayment(orderID, amount, dto.IdempotencyKey)
	if err != nil {
		return nil, err
	}

	// Save to repository
	err = uc.paymentRepo.CreatePayment(ctx, payment)
	if err != nil {
		return nil, err
	}

	return &application.CreatePaymentResponseDTO{
		ID:             payment.ID.String(),
		OrderID:        payment.OrderID.String(),
		Amount:         payment.Amount.String(),
		Status:         string(payment.Status),
		IdempotencyKey: payment.IdempotencyKey,
		CreatedAt:      payment.CreatedAt.String(),
	}, nil
}
