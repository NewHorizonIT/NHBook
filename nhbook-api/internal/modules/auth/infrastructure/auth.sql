-- name: UserExists :one
SELECT EXISTS (
        SELECT 1
        FROM users
        WHERE
            username = $1
            OR email = $2
    );

-- name: CreateUser :exec
INSERT INTO
    users (
        id,
        email,
        password,
        username,
        role,
        status
    )
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUserByEmail :one
SELECT id, email, password, username, role, status
FROM users
WHERE
    email = $1;