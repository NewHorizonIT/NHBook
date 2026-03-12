package infrastructure

import (
	"context"
	"database/sql"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/payment/infrastructure/sqlc"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func NewPaymentRepository(db *sql.DB) domain.IPaymentRepository {
	return &PaymentRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *domain.Payment) error {
	return r.q.CreatePayment(ctx, sqlc.CreatePaymentParams{
		ID:             payment.ID,
		OrderID:        uuid.NullUUID{UUID: payment.OrderID, Valid: true},
		IdempotencyKey: payment.IdempotencyKey,
		Amount:         payment.Amount.String(),
		Status:         string(payment.Status),
		CreatedAt:      sql.NullTime{Time: payment.CreatedAt, Valid: true},
		UpdatedAt:      sql.NullTime{Time: payment.UpdatedAt, Valid: true},
	})
}

func (r *PaymentRepository) GetPaymentByID(ctx context.Context, paymentID uuid.UUID) (*domain.Payment, error) {
	dbPayment, err := r.q.GetPaymentByID(ctx, paymentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, err
	}

	amount, _ := decimal.NewFromString(dbPayment.Amount)
	return &domain.Payment{
		ID:             dbPayment.ID,
		OrderID:        dbPayment.OrderID.UUID,
		IdempotencyKey: dbPayment.IdempotencyKey,
		Amount:         amount,
		Status:         domain.PaymentStatus(dbPayment.Status),
		CreatedAt:      dbPayment.CreatedAt.Time,
		UpdatedAt:      dbPayment.UpdatedAt.Time,
	}, nil
}

func (r *PaymentRepository) GetPaymentByIdempotencyKey(ctx context.Context, key string) (*domain.Payment, error) {
	dbPayment, err := r.q.GetPaymentByIdempotencyKey(ctx, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	amount, _ := decimal.NewFromString(dbPayment.Amount)
	return &domain.Payment{
		ID:             dbPayment.ID,
		OrderID:        dbPayment.OrderID.UUID,
		IdempotencyKey: dbPayment.IdempotencyKey,
		Amount:         amount,
		Status:         domain.PaymentStatus(dbPayment.Status),
		CreatedAt:      dbPayment.CreatedAt.Time,
		UpdatedAt:      dbPayment.UpdatedAt.Time,
	}, nil
}

func (r *PaymentRepository) GetPaymentByOrderID(ctx context.Context, orderID uuid.UUID) (*domain.Payment, error) {
	dbPayment, err := r.q.GetPaymentByOrderID(ctx, uuid.NullUUID{UUID: orderID, Valid: true})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, err
	}

	amount, _ := decimal.NewFromString(dbPayment.Amount)
	return &domain.Payment{
		ID:             dbPayment.ID,
		OrderID:        dbPayment.OrderID.UUID,
		IdempotencyKey: dbPayment.IdempotencyKey,
		Amount:         amount,
		Status:         domain.PaymentStatus(dbPayment.Status),
		CreatedAt:      dbPayment.CreatedAt.Time,
		UpdatedAt:      dbPayment.UpdatedAt.Time,
	}, nil
}

func (r *PaymentRepository) UpdatePaymentStatus(ctx context.Context, paymentID uuid.UUID, status domain.PaymentStatus) error {
	return r.q.UpdatePaymentStatus(ctx, sqlc.UpdatePaymentStatusParams{
		ID:     paymentID,
		Status: string(status),
	})
}
