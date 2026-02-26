-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    order_id UUID UNIQUE REFERENCES orders (id) ON DELETE CASCADE,
    amount NUMERIC(12, 2) NOT NULL,
    status VARCHAR(30) NOT NULL CHECK (
        status IN ('INIT', 'SUCCESS', 'FAILED')
    ),
    idempotency_key VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_payment_status ON payments (status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd