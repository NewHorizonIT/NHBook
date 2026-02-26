-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    bio TEXT,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE publishers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE books (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(12, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    author_id UUID REFERENCES authors (id),
    category_id UUID REFERENCES categories (id),
    publisher_id UUID REFERENCES publishers (id),
    cover_image TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_books_title ON books (title);

CREATE INDEX idx_books_author ON books (author_id);

CREATE INDEX idx_books_category ON books (category_id);

CREATE INDEX idx_books_publisher ON books (publisher_id);

CREATE INDEX idx_books_price ON books (price);

ALTER TABLE books ADD COLUMN search_vector tsvector;

CREATE INDEX idx_books_search ON books USING GIN (search_vector);

ALTER TABLE books ADD COLUMN search_vector tsvector;

CREATE INDEX idx_books_search ON books USING GIN (search_vector);

CREATE FUNCTION books_search_trigger() RETURNS trigger AS $$
BEGIN
  NEW.search_vector :=
    to_tsvector('simple',
      coalesce(NEW.title,'') || ' ' ||
      coalesce(NEW.description,''));
  RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdate
BEFORE INSERT OR UPDATE ON books
FOR EACH ROW EXECUTE FUNCTION books_search_trigger();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS books;

DROP TABLE IF EXISTS authors;

DROP TABLE IF EXISTS categories;
-- +goose StatementEnd