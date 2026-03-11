-- name: CreateOrder :exec
INSERT INTO
    orders (
        id,
        user_id,
        total_amount,
        status,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetOrderByID :one
SELECT
    id,
    user_id,
    total_amount,
    status,
    created_at,
    updated_at
FROM orders
WHERE
    id = $1;

-- name: GetOrdersByUserID :many
SELECT
    id,
    user_id,
    total_amount,
    status,
    created_at,
    updated_at
FROM orders
WHERE
    user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: GetOrdersCountByUserID :one
SELECT COUNT(*) as total FROM orders WHERE user_id = $1;

-- name: UpdateOrderStatus :exec
UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3;

-- name: GetOrderItems :many
SELECT
    id,
    order_id,
    book_id,
    book_title_snapshot,
    price_snapshot,
    quantity,
    created_at
FROM order_items
WHERE
    order_id = $1;

-- name: CreateOrderItem :exec
INSERT INTO
    order_items (
        id,
        order_id,
        book_id,
        book_title_snapshot,
        price_snapshot,
        quantity,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7);