package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/domain"
	"github.com/google/uuid"
)

// PaymentGateway defines interface for payment processing (e.g., Stripe, PayPal)
type PaymentGateway interface {
	ProcessPayment(ctx context.Context, paymentID uuid.UUID, amount string) (transactionID string, err error)
}

type ProcessPaymentUsecase struct {
	paymentRepo domain.IPaymentRepository
	gateway     PaymentGateway
}

func NewProcessPaymentUsecase(
	paymentRepo domain.IPaymentRepository,
	gateway PaymentGateway,
) *ProcessPaymentUsecase {
	return &ProcessPaymentUsecase{
		paymentRepo: paymentRepo,
		gateway:     gateway,
	}
}

func (uc *ProcessPaymentUsecase) Execute(
	ctx context.Context,
	dto *application.ProcessPaymentDTO,
) (*application.PaymentStatusDTO, error) {
	// Parse payment ID
	paymentID, err := uuid.Parse(dto.PaymentID)
	if err != nil {
		return nil, err
	}

	// Get payment
	payment, err := uc.paymentRepo.GetPaymentByID(ctx, paymentID)
	if err != nil {
		return nil, err
	}

	// Process payment through gateway
	_, err = uc.gateway.ProcessPayment(ctx, paymentID, payment.Amount.String())
	if err != nil {
		// Mark as failed
		_ = uc.paymentRepo.UpdatePaymentStatus(ctx, paymentID, domain.PaymentStatusFailed)
		payment.MarkAsFailed()
	} else {
		// Mark as success
		_ = uc.paymentRepo.UpdatePaymentStatus(ctx, paymentID, domain.PaymentStatusSuccess)
		payment.MarkAsSuccess()
	}

	return &application.PaymentStatusDTO{
		ID:        payment.ID.String(),
		OrderID:   payment.OrderID.String(),
		Status:    string(payment.Status),
		Amount:    payment.Amount.String(),
		UpdatedAt: payment.UpdatedAt.String(),
	}, nil
}
