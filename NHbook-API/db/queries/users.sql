-- name: CreateUser :execresult
INSERT INTO
    users (
        `id`,
        username,
        email,
        password_hash,
        phone
    )
VALUES (?, ?, ?, ?, ?);

-- name: GetUserByID :one
SELECT
    id,
    username,
    email,
    display_name,
    password_hash,
    phone,
    is_admin,
    created_at,
    updated_at
FROM users
WHERE
    id = ?
LIMIT 1;

-- name: GetUserByEmail :one
SELECT
    id,
    username,
    display_name,
    email,
    password_hash,
    phone,
    is_admin,
    created_at,
    updated_at
FROM users
WHERE
    email = ?
LIMIT 1;

-- name: ListUsers :many
SELECT
    id,
    username,
    display_name,
    email,
    phone,
    is_admin,
    created_at,
    updated_at
FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :exec
UPDATE users
SET
    display_name = COALESCE(?, full_name),
    email = COALESCE(?, email),
    phone = COALESCE(?, phone_number),
    is_admin = COALESCE(?, is_admin),
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = ?;