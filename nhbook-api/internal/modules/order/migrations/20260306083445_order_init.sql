-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID REFERENCES users (id),
    total_amount NUMERIC(12, 2) NOT NULL,
    status VARCHAR(30) NOT NULL CHECK (
        status IN (
            'PENDING',
            'PAID',
            'SHIPPED',
            'CANCELLED'
        )
    ),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_orders_user ON orders (user_id);

CREATE INDEX idx_orders_status ON orders (status);

CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    order_id UUID REFERENCES orders (id) ON DELETE CASCADE,
    book_id UUID REFERENCES books (id),
    book_title_snapshot VARCHAR(255),
    price_snapshot NUMERIC(12, 2),
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;

DROP TABLE IF EXISTS orders;
-- +goose StatementEnd