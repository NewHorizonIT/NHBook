package application

// Request DTOs
type CreatePaymentDTO struct {
	OrderID        string `json:"order_id" validate:"required"`
	IdempotencyKey string `json:"idempotency_key" validate:"required"`
}

type ProcessPaymentDTO struct {
	PaymentID string `json:"payment_id" validate:"required"`
}

// Response DTOs
type PaymentDTO struct {
	ID             string `json:"id"`
	OrderID        string `json:"order_id"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	IdempotencyKey string `json:"idempotency_key"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CreatePaymentResponseDTO struct {
	ID             string `json:"id"`
	OrderID        string `json:"order_id"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	IdempotencyKey string `json:"idempotency_key"`
	CreatedAt      string `json:"created_at"`
}

type PaymentStatusDTO struct {
	ID        string `json:"id"`
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	UpdatedAt string `json:"updated_at"`
}
