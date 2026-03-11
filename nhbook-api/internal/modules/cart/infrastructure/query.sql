-- name: GetCartByUserID :one
SELECT
    id,
    user_id,
    created_at,
    updated_at
FROM carts
WHERE
    user_id = $1;

-- name: CreateCart :exec
INSERT INTO
    carts (
        id,
        user_id,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4);

-- name: GetCartItems :many
SELECT
    id,
    cart_id,
    book_id,
    quantity,
    created_at,
    updated_at
FROM cart_items
WHERE
    cart_id = $1;

-- name: GetCartItemByBookID :one
SELECT
    id,
    cart_id,
    book_id,
    quantity,
    created_at,
    updated_at
FROM cart_items
WHERE
    cart_id = $1
    AND book_id = $2;

-- name: AddCartItem :exec
INSERT INTO
    cart_items (
        id,
        cart_id,
        book_id,
        quantity,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (cart_id, book_id) DO
UPDATE
SET
    quantity = cart_items.quantity + EXCLUDED.quantity,
    updated_at = EXCLUDED.updated_at;

-- name: UpdateCartItemQuantity :exec
UPDATE cart_items
SET
    quantity = $1,
    updated_at = $2
WHERE
    cart_id = $3
    AND book_id = $4;

-- name: RemoveCartItem :exec
DELETE FROM cart_items WHERE cart_id = $1 AND book_id = $2;

-- name: ClearCartItems :exec
DELETE FROM cart_items WHERE cart_id = $1;

-- name: DeleteCart :exec
DELETE FROM carts WHERE id = $1;