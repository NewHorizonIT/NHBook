-- ============================================================================
-- AUTHORS QUERIES
-- ============================================================================

-- name: CreateAuthor :one
INSERT INTO
    authors (id, name, bio)
VALUES ($1, $2, $3)
RETURNING
    *;

-- name: GetAuthorByID :one
SELECT * FROM authors WHERE id = $1;

-- name: GetAuthors :many
SELECT * FROM authors ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: GetAuthorsCount :one
SELECT COUNT(*) FROM authors;

-- name: UpdateAuthor :one
UPDATE authors
SET
    name = COALESCE(sqlc.narg ('name'), name),
    bio = COALESCE(sqlc.narg ('bio'), bio)
WHERE
    id = sqlc.arg ('id')
RETURNING
    *;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = $1;

-- name: AuthorExists :one
SELECT EXISTS ( SELECT 1 FROM authors WHERE id = $1 );

-- ============================================================================
-- CATEGORIES QUERIES
-- ============================================================================

-- name: CreateCategory :one
INSERT INTO
    categories (id, name)
VALUES ($1, $2)
RETURNING
    *;

-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = $1;

-- name: GetCategories :many
SELECT * FROM categories ORDER BY name ASC LIMIT $1 OFFSET $2;

-- name: GetCategoriesCount :one
SELECT COUNT(*) FROM categories;

-- name: GetAllCategories :many
SELECT * FROM categories ORDER BY name ASC;

-- name: UpdateCategory :one
UPDATE categories SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- name: CategoryExists :one
SELECT EXISTS ( SELECT 1 FROM categories WHERE id = $1 );

-- name: CategoryExistsByName :one
SELECT EXISTS ( SELECT 1 FROM categories WHERE name = $1 );

-- ============================================================================
-- PUBLISHERS QUERIES
-- ============================================================================

-- name: CreatePublisher :one
INSERT INTO
    publishers (id, name)
VALUES ($1, $2)
RETURNING
    *;

-- name: GetPublisherByID :one
SELECT * FROM publishers WHERE id = $1;

-- name: GetPublishers :many
SELECT * FROM publishers ORDER BY name ASC LIMIT $1 OFFSET $2;

-- name: GetPublishersCount :one
SELECT COUNT(*) FROM publishers;

-- name: GetAllPublishers :many
SELECT * FROM publishers ORDER BY name ASC;

-- name: UpdatePublisher :one
UPDATE publishers SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeletePublisher :exec
DELETE FROM publishers WHERE id = $1;

-- name: PublisherExists :one
SELECT EXISTS ( SELECT 1 FROM publishers WHERE id = $1 );

-- ============================================================================
-- BOOKS QUERIES
-- ============================================================================

-- name: CreateBook :one
INSERT INTO
    books (
        id,
        title,
        description,
        price,
        stock,
        author_id,
        category_id,
        publisher_id,
        cover_image,
        is_active
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    )
RETURNING
    *;

-- name: GetBookByID :one
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.id = $1;

-- name: GetBookByIDSimple :one
SELECT * FROM books WHERE id = $1;

-- name: GetBooks :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
ORDER BY b.created_at DESC
LIMIT $1
OFFSET
    $2;

-- name: GetBooksCount :one
SELECT COUNT(*) FROM books;

-- name: GetActiveBooks :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.is_active = true
ORDER BY b.created_at DESC
LIMIT $1
OFFSET
    $2;

-- name: GetActiveBooksCount :one
SELECT COUNT(*) FROM books WHERE is_active = true;

-- name: GetBooksByCategory :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.category_id = $1
    AND b.is_active = true
ORDER BY b.created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: GetBooksByAuthor :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.author_id = $1
    AND b.is_active = true
ORDER BY b.created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: GetBooksByPublisher :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.publisher_id = $1
    AND b.is_active = true
ORDER BY b.created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: SearchBooks :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name,
    ts_rank(
        b.search_vector,
        plainto_tsquery('simple', $1)
    ) as rank
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.search_vector @@ plainto_tsquery('simple', $1)
    AND b.is_active = true
ORDER BY rank DESC, b.created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: SearchBooksCount :one
SELECT COUNT(*)
FROM books b
WHERE
    b.search_vector @@ plainto_tsquery('simple', $1)
    AND b.is_active = true;

-- name: UpdateBook :one
UPDATE books
SET
    title = COALESCE(sqlc.narg ('title'), title),
    description = COALESCE(
        sqlc.narg ('description'),
        description
    ),
    price = COALESCE(sqlc.narg ('price'), price),
    stock = COALESCE(sqlc.narg ('stock'), stock),
    author_id = COALESCE(
        sqlc.narg ('author_id'),
        author_id
    ),
    category_id = COALESCE(
        sqlc.narg ('category_id'),
        category_id
    ),
    publisher_id = COALESCE(
        sqlc.narg ('publisher_id'),
        publisher_id
    ),
    cover_image = COALESCE(
        sqlc.narg ('cover_image'),
        cover_image
    ),
    updated_at = now()
WHERE
    id = sqlc.arg ('id')
RETURNING
    *;

-- name: DeleteBook :exec
DELETE FROM books WHERE id = $1;

-- name: AdjustStock :one
UPDATE books
SET
    stock = stock + $2,
    updated_at = now()
WHERE
    id = $1
RETURNING
    *;

-- name: ChangePrice :one
UPDATE books
SET
    price = $2,
    updated_at = now()
WHERE
    id = $1
RETURNING
    *;

-- name: ToggleBookActive :one
UPDATE books
SET
    is_active = NOT is_active,
    updated_at = now()
WHERE
    id = $1
RETURNING
    *;

-- name: SetBookActive :one
UPDATE books
SET
    is_active = $2,
    updated_at = now()
WHERE
    id = $1
RETURNING
    *;

-- name: BookExists :one
SELECT EXISTS ( SELECT 1 FROM books WHERE id = $1 );

-- name: CheckBookStock :one
SELECT stock FROM books WHERE id = $1;

-- name: GetBooksByPriceRange :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.price BETWEEN $1 AND $2
    AND b.is_active = true
ORDER BY b.price ASC
LIMIT $3
OFFSET
    $4;

-- name: GetLowStockBooks :many
SELECT
    b.*,
    a.name as author_name,
    c.name as category_name,
    p.name as publisher_name
FROM
    books b
    LEFT JOIN authors a ON b.author_id = a.id
    LEFT JOIN categories c ON b.category_id = c.id
    LEFT JOIN publishers p ON b.publisher_id = p.id
WHERE
    b.stock <= $1
ORDER BY b.stock ASC
LIMIT $2
OFFSET
    $3;