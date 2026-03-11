package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/domain"
	"github.com/google/uuid"
)

type GetPaymentStatusUsecase struct {
	repo domain.IPaymentRepository
}

func NewGetPaymentStatusUsecase(repo domain.IPaymentRepository) *GetPaymentStatusUsecase {
	return &GetPaymentStatusUsecase{
		repo: repo,
	}
}

func (uc *GetPaymentStatusUsecase) Execute(ctx context.Context, paymentID string) (*application.PaymentStatusDTO, error) {
	parsedPaymentID, err := uuid.Parse(paymentID)
	if err != nil {
		return nil, err
	}

	payment, err := uc.repo.GetPaymentByID(ctx, parsedPaymentID)
	if err != nil {
		return nil, err
	}

	return &application.PaymentStatusDTO{
		ID:        payment.ID.String(),
		OrderID:   payment.OrderID.String(),
		Status:    string(payment.Status),
		Amount:    payment.Amount.String(),
		UpdatedAt: payment.UpdatedAt.String(),
	}, nil
}
