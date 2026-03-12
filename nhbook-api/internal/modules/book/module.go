package book

import (
"database/sql"

authoruc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/author"
bookuc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/book"
categoryuc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/category"
publisheruc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/pulisher"
"github.com/NewHorizonIT/nhbook-api/internal/modules/book/infrastructure"
presentation "github.com/NewHorizonIT/nhbook-api/internal/modules/book/presentation"
"github.com/NewHorizonIT/nhbook-api/internal/shared/infrastructure/cache"
"github.com/gin-gonic/gin"
)

// BookModule quản lý toàn bộ book module
type BookModule struct {
	handler *presentation.BookHandler
}

// NewBookModule khởi tạo book module với tất cả dependencies
func NewBookModule(db *sql.DB, cacheService cache.ICache) *BookModule {
	// 1. Infrastructure Layer - Repositories
	authorRepo := infrastructure.NewAuthorRepository(db, cacheService)
	categoryRepo := infrastructure.NewCategoryRepository(db, cacheService)
	publisherRepo := infrastructure.NewPublisherRepository(db, cacheService)
	bookRepo := infrastructure.NewBookRepository(db, cacheService)

	// 2. Application Layer - Author UseCases
	createAuthorUsecase := authoruc.NewCreateAuthorUsecase(authorRepo)
	getAuthorByIDUsecase := authoruc.NewGetAuthorByIDUsecase(authorRepo)
	updateAuthorUsecase := authoruc.NewUpdateAuthorUsecase(authorRepo)
	deleteAuthorUsecase := authoruc.NewDeleteAuthorUsecase(authorRepo)

	// 3. Application Layer - Category UseCases
	createCategoryUsecase := categoryuc.NewCreateCategoryUsecase(categoryRepo)
	getCategoriesUsecase := categoryuc.NewGetCategoriesUsecase(categoryRepo)
	getCategoryByIDUsecase := categoryuc.NewGetCategoryByIDUsecase(categoryRepo)
	updateCategoryUsecase := categoryuc.NewUpdateCategoryUsecase(categoryRepo)
	deleteCategoryUsecase := categoryuc.NewDeleteCategoryUsecase(categoryRepo)

	// 4. Application Layer - Publisher UseCases
	createPublisherUsecase := publisheruc.NewCreatePublisherUsecase(publisherRepo)
	getPublishersUsecase := publisheruc.NewGetPublishersUsecase(publisherRepo)
	getPublisherByIDUsecase := publisheruc.NewGetPublisherByIDUsecase(publisherRepo)
	updatePublisherUsecase := publisheruc.NewUpdatePublisherUsecase(publisherRepo)
	deletePublisherUsecase := publisheruc.NewDeletePublisherUsecase(publisherRepo)

	// 5. Application Layer - Book UseCases
	createBookUsecase := bookuc.NewCreateBookUsecase(bookRepo, authorRepo, categoryRepo, publisherRepo)
	getBookByIDUsecase := bookuc.NewGetBookByIDUsecase(bookRepo)
	getBooksUsecase := bookuc.NewGetBooksUsecase(bookRepo)
	updateBookUsecase := bookuc.NewUpdateBookUsecase(bookRepo, authorRepo, categoryRepo, publisherRepo)
	deleteBookUsecase := bookuc.NewDeleteBookUsecase(bookRepo)
	searchBooksUsecase := bookuc.NewSearchBooksUsecase(bookRepo)
	adjustStockUsecase := bookuc.NewAdjustStockUsecase(bookRepo)
	changePriceUsecase := bookuc.NewChangePriceUsecase(bookRepo)
	toggleActiveUsecase := bookuc.NewToggleActiveUsecase(bookRepo)

	// 6. Presentation Layer - Handler
	handler := presentation.NewBookHandler(
// Author UseCases
createAuthorUsecase,
getAuthorByIDUsecase,
updateAuthorUsecase,
deleteAuthorUsecase,
// Category UseCases
createCategoryUsecase,
getCategoriesUsecase,
getCategoryByIDUsecase,
updateCategoryUsecase,
deleteCategoryUsecase,
// Publisher UseCases
createPublisherUsecase,
getPublishersUsecase,
getPublisherByIDUsecase,
updatePublisherUsecase,
deletePublisherUsecase,
// Book UseCases
createBookUsecase,
getBookByIDUsecase,
getBooksUsecase,
updateBookUsecase,
deleteBookUsecase,
searchBooksUsecase,
adjustStockUsecase,
changePriceUsecase,
toggleActiveUsecase,
)

	return &BookModule{
		handler: handler,
	}
}

// RegisterRoutes đăng ký routes cho module
func (m *BookModule) RegisterRoutes(router *gin.RouterGroup) {
	presentation.RegisterBookRoutes(router, m.handler)
}

// GetHandler trả về handler instance (nếu cần dùng ở nơi khác)
func (m *BookModule) GetHandler() *presentation.BookHandler {
	return m.handler
}
