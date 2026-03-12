-- name: CreatePayment :exec
INSERT INTO
    payments (
        id,
        order_id,
        amount,
        status,
        idempotency_key,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetPaymentByID :one
SELECT
    id,
    order_id,
    amount,
    status,
    idempotency_key,
    created_at,
    updated_at
FROM payments
WHERE
    id = $1;

-- name: GetPaymentByOrderID :one
SELECT
    id,
    order_id,
    amount,
    status,
    idempotency_key,
    created_at,
    updated_at
FROM payments
WHERE
    order_id = $1;

-- name: GetPaymentByIdempotencyKey :one
SELECT
    id,
    order_id,
    amount,
    status,
    idempotency_key,
    created_at,
    updated_at
FROM payments
WHERE
    idempotency_key = $1;

-- name: UpdatePaymentStatus :exec
UPDATE payments SET status = $1, updated_at = $2 WHERE id = $3;