package infrastructure

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/infrastructure/sqlc"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
	"github.com/NewHorizonIT/nhbook-api/internal/shared/infrastructure/cache"
	"github.com/NewHorizonIT/nhbook-api/pkg/errs"
	"github.com/NewHorizonIT/nhbook-api/pkg/helper"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Cache key prefixes
const (
	authorKeyPrefix    = "author:"
	categoryKeyPrefix  = "category:"
	publisherKeyPrefix = "publisher:"
	bookKeyPrefix      = "book:"

	// Cache TTLs
	defaultTTL     = 15 * time.Minute
	listTTL        = 5 * time.Minute
	shortTTL       = 2 * time.Minute
	entityCacheTTL = 30 * time.Minute
)

type AuthorRepository struct {
	db    *sql.DB
	q     *sqlc.Queries
	cache cache.ICache
}

func NewAuthorRepository(db *sql.DB, cache cache.ICache) domain.IAuthorRepository {
	return &AuthorRepository{
		db:    db,
		q:     sqlc.New(db),
		cache: cache,
	}
}

func (r *AuthorRepository) cacheKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", authorKeyPrefix, id.String())
}

func (r *AuthorRepository) Create(ctx context.Context, author *domain.Author) error {
	params := sqlc.CreateAuthorParams{
		ID:   author.ID,
		Name: author.Name,
		Bio:  helper.NullString(author.Bio),
	}

	_, err := r.q.CreateAuthor(ctx, params)
	if err != nil {
		return errs.Wrap(err, "AuthorRepository.Create", "INTERNAL_ERROR", "failed to create author")
	}

	// Cache the new author
	_ = r.cache.Set(ctx, r.cacheKey(author.ID), author, entityCacheTTL)

	return nil
}

func (r *AuthorRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Author, error) {
	cacheKey := r.cacheKey(id)

	// Try cache first (Cache-Aside: Read)
	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if author, ok := r.unmarshalAuthor(cached); ok {
			return author, nil
		}
	}

	// Cache miss - read from DB
	dbAuthor, err := r.q.GetAuthorByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrAuthorNotFound
		}
		return nil, errs.Wrap(err, "AuthorRepository.GetByID", "INTERNAL_ERROR", "failed to get author")
	}

	author := r.toDomainAuthor(dbAuthor)

	// Update cache
	_ = r.cache.Set(ctx, cacheKey, author, entityCacheTTL)

	return author, nil
}

func (r *AuthorRepository) GetAll(ctx context.Context, pagination shared.Pagination) ([]domain.Author, error) {
	params := sqlc.GetAuthorsParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	}

	dbAuthors, err := r.q.GetAuthors(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "AuthorRepository.GetAll", "INTERNAL_ERROR", "failed to get authors")
	}

	authors := make([]domain.Author, len(dbAuthors))
	for i, a := range dbAuthors {
		authors[i] = *r.toDomainAuthor(a)
	}

	return authors, nil
}

func (r *AuthorRepository) GetCount(ctx context.Context) (int64, error) {
	count, err := r.q.GetAuthorsCount(ctx)
	if err != nil {
		return 0, errs.Wrap(err, "AuthorRepository.GetCount", "INTERNAL_ERROR", "failed to get authors count")
	}
	return count, nil
}

func (r *AuthorRepository) Update(ctx context.Context, author *domain.Author) error {
	params := sqlc.UpdateAuthorParams{
		ID:   author.ID,
		Name: sql.NullString{String: author.Name, Valid: true},
		Bio:  helper.NullString(author.Bio),
	}

	_, err := r.q.UpdateAuthor(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrAuthorNotFound
		}
		return errs.Wrap(err, "AuthorRepository.Update", "INTERNAL_ERROR", "failed to update author")
	}

	// Invalidate cache (Cache-Aside: Write)
	_ = r.cache.Delete(ctx, r.cacheKey(author.ID))

	return nil
}

func (r *AuthorRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteAuthor(ctx, id)
	if err != nil {
		return errs.Wrap(err, "AuthorRepository.Delete", "INTERNAL_ERROR", "failed to delete author")
	}

	// Invalidate cache
	_ = r.cache.Delete(ctx, r.cacheKey(id))

	return nil
}

func (r *AuthorRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	exists, err := r.q.AuthorExists(ctx, id)
	if err != nil {
		return false, errs.Wrap(err, "AuthorRepository.Exists", "INTERNAL_ERROR", "failed to check author exists")
	}
	return exists, nil
}

func (r *AuthorRepository) toDomainAuthor(a sqlc.Author) *domain.Author {
	return &domain.Author{
		ID:        a.ID,
		Name:      a.Name,
		Bio:       helper.PtrString(a.Bio),
		CreatedAt: helper.PtrTime(a.CreatedAt),
	}
}

func (r *AuthorRepository) unmarshalAuthor(data any) (*domain.Author, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var author domain.Author
	if err := json.Unmarshal(bytes, &author); err != nil {
		return nil, false
	}
	return &author, true
}

// ============================================================================
// CATEGORY REPOSITORY
// ============================================================================

type CategoryRepository struct {
	db    *sql.DB
	q     *sqlc.Queries
	cache cache.ICache
}

func NewCategoryRepository(db *sql.DB, cache cache.ICache) domain.ICategoryRepository {
	return &CategoryRepository{
		db:    db,
		q:     sqlc.New(db),
		cache: cache,
	}
}

func (r *CategoryRepository) cacheKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", categoryKeyPrefix, id.String())
}

func (r *CategoryRepository) allCacheKey() string {
	return categoryKeyPrefix + "all"
}

func (r *CategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	params := sqlc.CreateCategoryParams{
		ID:   category.ID,
		Name: category.Name,
	}

	_, err := r.q.CreateCategory(ctx, params)
	if err != nil {
		return errs.Wrap(err, "CategoryRepository.Create", "INTERNAL_ERROR", "failed to create category")
	}

	// Cache the new category & invalidate list cache
	_ = r.cache.Set(ctx, r.cacheKey(category.ID), category, entityCacheTTL)
	_ = r.cache.Delete(ctx, r.allCacheKey())

	return nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	cacheKey := r.cacheKey(id)

	// Try cache first
	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if category, ok := r.unmarshalCategory(cached); ok {
			return category, nil
		}
	}

	// Cache miss
	dbCategory, err := r.q.GetCategoryByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, errs.Wrap(err, "CategoryRepository.GetByID", "INTERNAL_ERROR", "failed to get category")
	}

	category := r.toDomainCategory(dbCategory)
	_ = r.cache.Set(ctx, cacheKey, category, entityCacheTTL)

	return category, nil
}

func (r *CategoryRepository) GetAll(ctx context.Context, pagination shared.Pagination) ([]domain.Category, error) {
	params := sqlc.GetCategoriesParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	}

	dbCategories, err := r.q.GetCategories(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "CategoryRepository.GetAll", "INTERNAL_ERROR", "failed to get categories")
	}

	categories := make([]domain.Category, len(dbCategories))
	for i, c := range dbCategories {
		categories[i] = *r.toDomainCategory(c)
	}

	return categories, nil
}

func (r *CategoryRepository) GetAllWithoutPagination(ctx context.Context) ([]domain.Category, error) {
	cacheKey := r.allCacheKey()

	// Try cache first
	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if categories, ok := r.unmarshalCategories(cached); ok {
			return categories, nil
		}
	}

	dbCategories, err := r.q.GetAllCategories(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "CategoryRepository.GetAllWithoutPagination", "INTERNAL_ERROR", "failed to get all categories")
	}

	categories := make([]domain.Category, len(dbCategories))
	for i, c := range dbCategories {
		categories[i] = *r.toDomainCategory(c)
	}

	_ = r.cache.Set(ctx, cacheKey, categories, listTTL)

	return categories, nil
}

func (r *CategoryRepository) GetCount(ctx context.Context) (int64, error) {
	count, err := r.q.GetCategoriesCount(ctx)
	if err != nil {
		return 0, errs.Wrap(err, "CategoryRepository.GetCount", "INTERNAL_ERROR", "failed to get categories count")
	}
	return count, nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *domain.Category) error {
	params := sqlc.UpdateCategoryParams{
		ID:   category.ID,
		Name: category.Name,
	}

	_, err := r.q.UpdateCategory(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrCategoryNotFound
		}
		return errs.Wrap(err, "CategoryRepository.Update", "INTERNAL_ERROR", "failed to update category")
	}

	// Invalidate caches
	_ = r.cache.Delete(ctx, r.cacheKey(category.ID))
	_ = r.cache.Delete(ctx, r.allCacheKey())

	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteCategory(ctx, id)
	if err != nil {
		return errs.Wrap(err, "CategoryRepository.Delete", "INTERNAL_ERROR", "failed to delete category")
	}

	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.allCacheKey())

	return nil
}

func (r *CategoryRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	exists, err := r.q.CategoryExists(ctx, id)
	if err != nil {
		return false, errs.Wrap(err, "CategoryRepository.Exists", "INTERNAL_ERROR", "failed to check category exists")
	}
	return exists, nil
}

func (r *CategoryRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	exists, err := r.q.CategoryExistsByName(ctx, name)
	if err != nil {
		return false, errs.Wrap(err, "CategoryRepository.ExistsByName", "INTERNAL_ERROR", "failed to check category name exists")
	}
	return exists, nil
}

func (r *CategoryRepository) toDomainCategory(c sqlc.Category) *domain.Category {
	return &domain.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: helper.PtrTime(c.CreatedAt),
	}
}

func (r *CategoryRepository) unmarshalCategory(data any) (*domain.Category, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var category domain.Category
	if err := json.Unmarshal(bytes, &category); err != nil {
		return nil, false
	}
	return &category, true
}

func (r *CategoryRepository) unmarshalCategories(data any) ([]domain.Category, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var categories []domain.Category
	if err := json.Unmarshal(bytes, &categories); err != nil {
		return nil, false
	}
	return categories, true
}

// ============================================================================
// PUBLISHER REPOSITORY
// ============================================================================

type PublisherRepository struct {
	db    *sql.DB
	q     *sqlc.Queries
	cache cache.ICache
}

func NewPublisherRepository(db *sql.DB, cache cache.ICache) domain.IPublisherRepository {
	return &PublisherRepository{
		db:    db,
		q:     sqlc.New(db),
		cache: cache,
	}
}

func (r *PublisherRepository) cacheKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", publisherKeyPrefix, id.String())
}

func (r *PublisherRepository) allCacheKey() string {
	return publisherKeyPrefix + "all"
}

func (r *PublisherRepository) Create(ctx context.Context, publisher *domain.Publisher) error {
	params := sqlc.CreatePublisherParams{
		ID:   publisher.ID,
		Name: publisher.Name,
	}

	_, err := r.q.CreatePublisher(ctx, params)
	if err != nil {
		return errs.Wrap(err, "PublisherRepository.Create", "INTERNAL_ERROR", "failed to create publisher")
	}

	_ = r.cache.Set(ctx, r.cacheKey(publisher.ID), publisher, entityCacheTTL)
	_ = r.cache.Delete(ctx, r.allCacheKey())

	return nil
}

func (r *PublisherRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Publisher, error) {
	cacheKey := r.cacheKey(id)

	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if publisher, ok := r.unmarshalPublisher(cached); ok {
			return publisher, nil
		}
	}

	dbPublisher, err := r.q.GetPublisherByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrPublisherNotFound
		}
		return nil, errs.Wrap(err, "PublisherRepository.GetByID", "INTERNAL_ERROR", "failed to get publisher")
	}

	publisher := r.toDomainPublisher(dbPublisher)
	_ = r.cache.Set(ctx, cacheKey, publisher, entityCacheTTL)

	return publisher, nil
}

func (r *PublisherRepository) GetAll(ctx context.Context, pagination shared.Pagination) ([]domain.Publisher, error) {
	params := sqlc.GetPublishersParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	}

	dbPublishers, err := r.q.GetPublishers(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "PublisherRepository.GetAll", "INTERNAL_ERROR", "failed to get publishers")
	}

	publishers := make([]domain.Publisher, len(dbPublishers))
	for i, p := range dbPublishers {
		publishers[i] = *r.toDomainPublisher(p)
	}

	return publishers, nil
}

func (r *PublisherRepository) GetAllWithoutPagination(ctx context.Context) ([]domain.Publisher, error) {
	cacheKey := r.allCacheKey()

	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if publishers, ok := r.unmarshalPublishers(cached); ok {
			return publishers, nil
		}
	}

	dbPublishers, err := r.q.GetAllPublishers(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "PublisherRepository.GetAllWithoutPagination", "INTERNAL_ERROR", "failed to get all publishers")
	}

	publishers := make([]domain.Publisher, len(dbPublishers))
	for i, p := range dbPublishers {
		publishers[i] = *r.toDomainPublisher(p)
	}

	_ = r.cache.Set(ctx, cacheKey, publishers, listTTL)

	return publishers, nil
}

func (r *PublisherRepository) GetCount(ctx context.Context) (int64, error) {
	count, err := r.q.GetPublishersCount(ctx)
	if err != nil {
		return 0, errs.Wrap(err, "PublisherRepository.GetCount", "INTERNAL_ERROR", "failed to get publishers count")
	}
	return count, nil
}

func (r *PublisherRepository) Update(ctx context.Context, publisher *domain.Publisher) error {
	params := sqlc.UpdatePublisherParams{
		ID:   publisher.ID,
		Name: publisher.Name,
	}

	_, err := r.q.UpdatePublisher(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrPublisherNotFound
		}
		return errs.Wrap(err, "PublisherRepository.Update", "INTERNAL_ERROR", "failed to update publisher")
	}

	_ = r.cache.Delete(ctx, r.cacheKey(publisher.ID))
	_ = r.cache.Delete(ctx, r.allCacheKey())

	return nil
}

func (r *PublisherRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeletePublisher(ctx, id)
	if err != nil {
		return errs.Wrap(err, "PublisherRepository.Delete", "INTERNAL_ERROR", "failed to delete publisher")
	}

	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.allCacheKey())

	return nil
}

func (r *PublisherRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	exists, err := r.q.PublisherExists(ctx, id)
	if err != nil {
		return false, errs.Wrap(err, "PublisherRepository.Exists", "INTERNAL_ERROR", "failed to check publisher exists")
	}
	return exists, nil
}

func (r *PublisherRepository) toDomainPublisher(p sqlc.Publisher) *domain.Publisher {
	return &domain.Publisher{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: helper.PtrTime(p.CreatedAt),
	}
}

func (r *PublisherRepository) unmarshalPublisher(data any) (*domain.Publisher, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var publisher domain.Publisher
	if err := json.Unmarshal(bytes, &publisher); err != nil {
		return nil, false
	}
	return &publisher, true
}

func (r *PublisherRepository) unmarshalPublishers(data any) ([]domain.Publisher, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var publishers []domain.Publisher
	if err := json.Unmarshal(bytes, &publishers); err != nil {
		return nil, false
	}
	return publishers, true
}

// ============================================================================
// BOOK REPOSITORY
// ============================================================================

type BookRepository struct {
	db    *sql.DB
	q     *sqlc.Queries
	cache cache.ICache
}

func NewBookRepository(db *sql.DB, cache cache.ICache) domain.IBookRepository {
	return &BookRepository{
		db:    db,
		q:     sqlc.New(db),
		cache: cache,
	}
}

func (r *BookRepository) cacheKey(id uuid.UUID) string {
	return fmt.Sprintf("%s%s", bookKeyPrefix, id.String())
}

func (r *BookRepository) cacheKeyWithRelations(id uuid.UUID) string {
	return fmt.Sprintf("%srelations:%s", bookKeyPrefix, id.String())
}

// CRUD Operations

func (r *BookRepository) Create(ctx context.Context, book *domain.Book) error {
	params := sqlc.CreateBookParams{
		ID:          book.ID,
		Title:       book.Title,
		Description: helper.NullString(book.Description),
		Price:       book.Price.String(),
		Stock:       book.Stock.Value(),
		AuthorID:    helper.NullUUID(book.AuthorID),
		CategoryID:  helper.NullUUID(book.CategoryID),
		PublisherID: helper.NullUUID(book.PublisherID),
		CoverImage:  helper.NullString(book.CoverImage),
		IsActive:    helper.NullBool(book.IsActive),
	}

	_, err := r.q.CreateBook(ctx, params)
	if err != nil {
		return errs.Wrap(err, "BookRepository.Create", "INTERNAL_ERROR", "failed to create book")
	}

	_ = r.cache.Set(ctx, r.cacheKey(book.ID), book, entityCacheTTL)

	return nil
}

func (r *BookRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Book, error) {
	cacheKey := r.cacheKey(id)

	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if book, ok := r.unmarshalBook(cached); ok {
			return book, nil
		}
	}

	dbBook, err := r.q.GetBookByIDSimple(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, errs.Wrap(err, "BookRepository.GetByID", "INTERNAL_ERROR", "failed to get book")
	}

	book := r.toDomainBook(dbBook)
	_ = r.cache.Set(ctx, cacheKey, book, entityCacheTTL)

	return book, nil
}

func (r *BookRepository) GetByIDWithRelations(ctx context.Context, id uuid.UUID) (*domain.BookWithRelations, error) {
	cacheKey := r.cacheKeyWithRelations(id)

	if cached, err := r.cache.Get(ctx, cacheKey); err == nil {
		if book, ok := r.unmarshalBookWithRelations(cached); ok {
			return book, nil
		}
	}

	dbBook, err := r.q.GetBookByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, errs.Wrap(err, "BookRepository.GetByIDWithRelations", "INTERNAL_ERROR", "failed to get book")
	}

	book := r.rowToBookWithRelations(dbBook)
	_ = r.cache.Set(ctx, cacheKey, book, defaultTTL)

	return book, nil
}

func (r *BookRepository) Update(ctx context.Context, book *domain.Book) error {
	params := sqlc.UpdateBookParams{
		ID:          book.ID,
		Title:       sql.NullString{String: book.Title, Valid: true},
		Description: helper.NullString(book.Description),
		Price:       sql.NullString{String: book.Price.String(), Valid: true},
		Stock:       sql.NullInt32{Int32: book.Stock.Value(), Valid: true},
		AuthorID:    helper.NullUUID(book.AuthorID),
		CategoryID:  helper.NullUUID(book.CategoryID),
		PublisherID: helper.NullUUID(book.PublisherID),
		CoverImage:  helper.NullString(book.CoverImage),
	}

	_, err := r.q.UpdateBook(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrBookNotFound
		}
		return errs.Wrap(err, "BookRepository.Update", "INTERNAL_ERROR", "failed to update book")
	}

	// Invalidate caches
	_ = r.cache.Delete(ctx, r.cacheKey(book.ID))
	_ = r.cache.Delete(ctx, r.cacheKeyWithRelations(book.ID))

	return nil
}

func (r *BookRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.q.DeleteBook(ctx, id)
	if err != nil {
		return errs.Wrap(err, "BookRepository.Delete", "INTERNAL_ERROR", "failed to delete book")
	}

	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.cacheKeyWithRelations(id))

	return nil
}

func (r *BookRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	exists, err := r.q.BookExists(ctx, id)
	if err != nil {
		return false, errs.Wrap(err, "BookRepository.Exists", "INTERNAL_ERROR", "failed to check book exists")
	}
	return exists, nil
}

// List & Filter Operations

func (r *BookRepository) GetAll(ctx context.Context, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetBooksParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	}

	rows, err := r.q.GetBooks(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetAll", "INTERNAL_ERROR", "failed to get books")
	}

	return r.getBooksRowsToBookWithRelations(rows), nil
}

func (r *BookRepository) GetCount(ctx context.Context) (int64, error) {
	count, err := r.q.GetBooksCount(ctx)
	if err != nil {
		return 0, errs.Wrap(err, "BookRepository.GetCount", "INTERNAL_ERROR", "failed to get books count")
	}
	return count, nil
}

func (r *BookRepository) GetActive(ctx context.Context, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetActiveBooksParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	}

	rows, err := r.q.GetActiveBooks(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetActive", "INTERNAL_ERROR", "failed to get active books")
	}

	return r.getActiveBooksRowsToBookWithRelations(rows), nil
}

func (r *BookRepository) GetActiveCount(ctx context.Context) (int64, error) {
	count, err := r.q.GetActiveBooksCount(ctx)
	if err != nil {
		return 0, errs.Wrap(err, "BookRepository.GetActiveCount", "INTERNAL_ERROR", "failed to get active books count")
	}
	return count, nil
}

func (r *BookRepository) GetByCategory(ctx context.Context, categoryID uuid.UUID, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetBooksByCategoryParams{
		CategoryID: uuid.NullUUID{UUID: categoryID, Valid: true},
		Limit:      pagination.Limit(),
		Offset:     pagination.Offset(),
	}

	rows, err := r.q.GetBooksByCategory(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetByCategory", "INTERNAL_ERROR", "failed to get books by category")
	}

	return r.getBooksByCategoryRowsToBookWithRelations(rows), nil
}

func (r *BookRepository) GetByAuthor(ctx context.Context, authorID uuid.UUID, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetBooksByAuthorParams{
		AuthorID: uuid.NullUUID{UUID: authorID, Valid: true},
		Limit:    pagination.Limit(),
		Offset:   pagination.Offset(),
	}

	rows, err := r.q.GetBooksByAuthor(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetByAuthor", "INTERNAL_ERROR", "failed to get books by author")
	}

	return r.getBooksByAuthorRowsToBookWithRelations(rows), nil
}

func (r *BookRepository) GetByPublisher(ctx context.Context, publisherID uuid.UUID, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetBooksByPublisherParams{
		PublisherID: uuid.NullUUID{UUID: publisherID, Valid: true},
		Limit:       pagination.Limit(),
		Offset:      pagination.Offset(),
	}

	rows, err := r.q.GetBooksByPublisher(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetByPublisher", "INTERNAL_ERROR", "failed to get books by publisher")
	}

	return r.getBooksByPublisherRowsToBookWithRelations(rows), nil
}

func (r *BookRepository) GetByPriceRange(ctx context.Context, minPrice, maxPrice decimal.Decimal, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetBooksByPriceRangeParams{
		Price:   minPrice.StringFixed(2),
		Price_2: maxPrice.StringFixed(2),
		Limit:   pagination.Limit(),
		Offset:  pagination.Offset(),
	}

	rows, err := r.q.GetBooksByPriceRange(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetByPriceRange", "INTERNAL_ERROR", "failed to get books by price range")
	}

	return r.getBooksByPriceRangeRowsToBookWithRelations(rows), nil
}

// Search Operations

func (r *BookRepository) Search(ctx context.Context, query string, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.SearchBooksParams{
		PlaintoTsquery: query,
		Limit:          pagination.Limit(),
		Offset:         pagination.Offset(),
	}

	rows, err := r.q.SearchBooks(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.Search", "INTERNAL_ERROR", "failed to search books")
	}

	return r.searchBooksRowsToBookWithRelations(rows), nil
}

func (r *BookRepository) SearchCount(ctx context.Context, query string) (int64, error) {
	count, err := r.q.SearchBooksCount(ctx, query)
	if err != nil {
		return 0, errs.Wrap(err, "BookRepository.SearchCount", "INTERNAL_ERROR", "failed to get search count")
	}
	return count, nil
}

// Stock Management

func (r *BookRepository) AdjustStock(ctx context.Context, id uuid.UUID, delta int32) (*domain.Book, error) {
	params := sqlc.AdjustStockParams{
		ID:    id,
		Stock: delta,
	}

	dbBook, err := r.q.AdjustStock(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, errs.Wrap(err, "BookRepository.AdjustStock", "INTERNAL_ERROR", "failed to adjust stock")
	}

	book := r.toDomainBook(dbBook)

	// Update cache
	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.cacheKeyWithRelations(id))

	return book, nil
}

func (r *BookRepository) CheckStock(ctx context.Context, id uuid.UUID) (int32, error) {
	stock, err := r.q.CheckBookStock(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, domain.ErrBookNotFound
		}
		return 0, errs.Wrap(err, "BookRepository.CheckStock", "INTERNAL_ERROR", "failed to check stock")
	}
	return stock, nil
}

func (r *BookRepository) GetLowStock(ctx context.Context, threshold int32, pagination shared.Pagination) ([]domain.BookWithRelations, error) {
	params := sqlc.GetLowStockBooksParams{
		Stock:  threshold,
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	}

	rows, err := r.q.GetLowStockBooks(ctx, params)
	if err != nil {
		return nil, errs.Wrap(err, "BookRepository.GetLowStock", "INTERNAL_ERROR", "failed to get low stock books")
	}

	return r.getLowStockBooksRowsToBookWithRelations(rows), nil
}

// Price Management

func (r *BookRepository) ChangePrice(ctx context.Context, id uuid.UUID, newPrice decimal.Decimal) (*domain.Book, error) {
	params := sqlc.ChangePriceParams{
		ID:    id,
		Price: newPrice.StringFixed(2),
	}

	dbBook, err := r.q.ChangePrice(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, errs.Wrap(err, "BookRepository.ChangePrice", "INTERNAL_ERROR", "failed to change price")
	}

	book := r.toDomainBook(dbBook)

	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.cacheKeyWithRelations(id))

	return book, nil
}

// Status Management

func (r *BookRepository) ToggleActive(ctx context.Context, id uuid.UUID) (*domain.Book, error) {
	dbBook, err := r.q.ToggleBookActive(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, errs.Wrap(err, "BookRepository.ToggleActive", "INTERNAL_ERROR", "failed to toggle active status")
	}

	book := r.toDomainBook(dbBook)

	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.cacheKeyWithRelations(id))

	return book, nil
}

func (r *BookRepository) SetActive(ctx context.Context, id uuid.UUID, isActive bool) (*domain.Book, error) {
	params := sqlc.SetBookActiveParams{
		ID:       id,
		IsActive: helper.NullBool(isActive),
	}

	dbBook, err := r.q.SetBookActive(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, errs.Wrap(err, "BookRepository.SetActive", "INTERNAL_ERROR", "failed to set active status")
	}

	book := r.toDomainBook(dbBook)

	_ = r.cache.Delete(ctx, r.cacheKey(id))
	_ = r.cache.Delete(ctx, r.cacheKeyWithRelations(id))

	return book, nil
}

// ============================================================================
// BOOK HELPER METHODS
// ============================================================================

func (r *BookRepository) toDomainBook(b sqlc.Book) *domain.Book {
	price, _ := domain.NewPriceFromString(b.Price)
	stock, _ := domain.NewStock(b.Stock)

	return &domain.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: helper.PtrString(b.Description),
		Price:       price,
		Stock:       stock,
		AuthorID:    helper.PtrUUID(b.AuthorID),
		CategoryID:  helper.PtrUUID(b.CategoryID),
		PublisherID: helper.PtrUUID(b.PublisherID),
		CoverImage:  helper.PtrString(b.CoverImage),
		IsActive:    helper.PtrBool(b.IsActive),
		CreatedAt:   helper.PtrTime(b.CreatedAt),
		UpdatedAt:   helper.PtrTime(b.UpdatedAt),
	}
}

func (r *BookRepository) rowToBookWithRelations(row sqlc.GetBookByIDRow) *domain.BookWithRelations {
	price, _ := domain.NewPriceFromString(row.Price)
	stock, _ := domain.NewStock(row.Stock)

	return &domain.BookWithRelations{
		Book: domain.Book{
			ID:          row.ID,
			Title:       row.Title,
			Description: helper.PtrString(row.Description),
			Price:       price,
			Stock:       stock,
			AuthorID:    helper.PtrUUID(row.AuthorID),
			CategoryID:  helper.PtrUUID(row.CategoryID),
			PublisherID: helper.PtrUUID(row.PublisherID),
			CoverImage:  helper.PtrString(row.CoverImage),
			IsActive:    helper.PtrBool(row.IsActive),
			CreatedAt:   helper.PtrTime(row.CreatedAt),
			UpdatedAt:   helper.PtrTime(row.UpdatedAt),
		},
		AuthorName:    helper.PtrString(row.AuthorName),
		CategoryName:  helper.PtrString(row.CategoryName),
		PublisherName: helper.PtrString(row.PublisherName),
	}
}

func (r *BookRepository) getBooksRowsToBookWithRelations(rows []sqlc.GetBooksRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) getActiveBooksRowsToBookWithRelations(rows []sqlc.GetActiveBooksRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) getBooksByCategoryRowsToBookWithRelations(rows []sqlc.GetBooksByCategoryRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) getBooksByAuthorRowsToBookWithRelations(rows []sqlc.GetBooksByAuthorRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) getBooksByPublisherRowsToBookWithRelations(rows []sqlc.GetBooksByPublisherRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) getBooksByPriceRangeRowsToBookWithRelations(rows []sqlc.GetBooksByPriceRangeRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) searchBooksRowsToBookWithRelations(rows []sqlc.SearchBooksRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) getLowStockBooksRowsToBookWithRelations(rows []sqlc.GetLowStockBooksRow) []domain.BookWithRelations {
	books := make([]domain.BookWithRelations, len(rows))
	for i, row := range rows {
		price, _ := domain.NewPriceFromString(row.Price)
		stock, _ := domain.NewStock(row.Stock)
		books[i] = domain.BookWithRelations{
			Book: domain.Book{
				ID:          row.ID,
				Title:       row.Title,
				Description: helper.PtrString(row.Description),
				Price:       price,
				Stock:       stock,
				AuthorID:    helper.PtrUUID(row.AuthorID),
				CategoryID:  helper.PtrUUID(row.CategoryID),
				PublisherID: helper.PtrUUID(row.PublisherID),
				CoverImage:  helper.PtrString(row.CoverImage),
				IsActive:    helper.PtrBool(row.IsActive),
				CreatedAt:   helper.PtrTime(row.CreatedAt),
				UpdatedAt:   helper.PtrTime(row.UpdatedAt),
			},
			AuthorName:    helper.PtrString(row.AuthorName),
			CategoryName:  helper.PtrString(row.CategoryName),
			PublisherName: helper.PtrString(row.PublisherName),
		}
	}
	return books
}

func (r *BookRepository) unmarshalBook(data any) (*domain.Book, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var book domain.Book
	if err := json.Unmarshal(bytes, &book); err != nil {
		return nil, false
	}
	return &book, true
}

func (r *BookRepository) unmarshalBookWithRelations(data any) (*domain.BookWithRelations, bool) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, false
	}
	var book domain.BookWithRelations
	if err := json.Unmarshal(bytes, &book); err != nil {
		return nil, false
	}
	return &book, true
}
