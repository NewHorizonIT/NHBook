package presentation

import (
"net/http"

"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
authoruc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/author"
bookuc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/book"
categoryuc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/category"
publisheruc "github.com/NewHorizonIT/nhbook-api/internal/modules/book/application/usecase/pulisher"
"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
"github.com/NewHorizonIT/nhbook-api/pkg/errs"
"github.com/NewHorizonIT/nhbook-api/pkg/response"
"github.com/NewHorizonIT/nhbook-api/pkg/validator"
"github.com/gin-gonic/gin"
validatorpkg "github.com/go-playground/validator/v10"
"github.com/google/uuid"
)

// BookHandler chứa tất cả dependencies cần thiết
type BookHandler struct {
	// Author UseCases
	createAuthorUsecase    *authoruc.CreateAuthorUsecase
	getAuthorByIDUsecase   *authoruc.GetAuthorByIDUsecase
	updateAuthorUsecase    *authoruc.UpdateAuthorUsecase
	deleteAuthorUsecase    *authoruc.DeleteAuthorUsecase
	// Category UseCases
	createCategoryUsecase    *categoryuc.CreateCategoryUsecase
	getCategoriesUsecase     *categoryuc.GetCategoriesUsecase
	getCategoryByIDUsecase   *categoryuc.GetCategoryByIDUsecase
	updateCategoryUsecase    *categoryuc.UpdateCategoryUsecase
	deleteCategoryUsecase    *categoryuc.DeleteCategoryUsecase
	// Publisher UseCases
	createPublisherUsecase    *publisheruc.CreatePublisherUsecase
	getPublishersUsecase      *publisheruc.GetPublishersUsecase
	getPublisherByIDUsecase   *publisheruc.GetPublisherByIDUsecase
	updatePublisherUsecase    *publisheruc.UpdatePublisherUsecase
	deletePublisherUsecase    *publisheruc.DeletePublisherUsecase
	// Book UseCases
	createBookUsecase      *bookuc.CreateBookUsecase
	getBookByIDUsecase     *bookuc.GetBookByIDUsecase
	getBooksUsecase        *bookuc.GetBooksUsecase
	updateBookUsecase      *bookuc.UpdateBookUsecase
	deleteBookUsecase      *bookuc.DeleteBookUsecase
	searchBooksUsecase     *bookuc.SearchBooksUsecase
	adjustStockUsecase     *bookuc.AdjustStockUsecase
	changePriceUsecase     *bookuc.ChangePriceUsecase
	toggleActiveUsecase    *bookuc.ToggleActiveUsecase
	validator              *validatorpkg.Validate
}

// NewBookHandler khởi tạo handler với dependency injection
func NewBookHandler(
createAuthorUsecase *authoruc.CreateAuthorUsecase,
getAuthorByIDUsecase *authoruc.GetAuthorByIDUsecase,
updateAuthorUsecase *authoruc.UpdateAuthorUsecase,
deleteAuthorUsecase *authoruc.DeleteAuthorUsecase,
createCategoryUsecase *categoryuc.CreateCategoryUsecase,
getCategoriesUsecase *categoryuc.GetCategoriesUsecase,
getCategoryByIDUsecase *categoryuc.GetCategoryByIDUsecase,
updateCategoryUsecase *categoryuc.UpdateCategoryUsecase,
deleteCategoryUsecase *categoryuc.DeleteCategoryUsecase,
createPublisherUsecase *publisheruc.CreatePublisherUsecase,
getPublishersUsecase *publisheruc.GetPublishersUsecase,
getPublisherByIDUsecase *publisheruc.GetPublisherByIDUsecase,
updatePublisherUsecase *publisheruc.UpdatePublisherUsecase,
deletePublisherUsecase *publisheruc.DeletePublisherUsecase,
createBookUsecase *bookuc.CreateBookUsecase,
getBookByIDUsecase *bookuc.GetBookByIDUsecase,
getBooksUsecase *bookuc.GetBooksUsecase,
updateBookUsecase *bookuc.UpdateBookUsecase,
deleteBookUsecase *bookuc.DeleteBookUsecase,
searchBooksUsecase *bookuc.SearchBooksUsecase,
adjustStockUsecase *bookuc.AdjustStockUsecase,
changePriceUsecase *bookuc.ChangePriceUsecase,
toggleActiveUsecase *bookuc.ToggleActiveUsecase,
) *BookHandler {
	return &BookHandler{
		createAuthorUsecase:    createAuthorUsecase,
		getAuthorByIDUsecase:   getAuthorByIDUsecase,
		updateAuthorUsecase:    updateAuthorUsecase,
		deleteAuthorUsecase:    deleteAuthorUsecase,
		createCategoryUsecase:  createCategoryUsecase,
		getCategoriesUsecase:   getCategoriesUsecase,
		getCategoryByIDUsecase: getCategoryByIDUsecase,
		updateCategoryUsecase:  updateCategoryUsecase,
		deleteCategoryUsecase:  deleteCategoryUsecase,
		createPublisherUsecase: createPublisherUsecase,
		getPublishersUsecase:   getPublishersUsecase,
		getPublisherByIDUsecase: getPublisherByIDUsecase,
		updatePublisherUsecase: updatePublisherUsecase,
		deletePublisherUsecase: deletePublisherUsecase,
		createBookUsecase:      createBookUsecase,
		getBookByIDUsecase:     getBookByIDUsecase,
		getBooksUsecase:        getBooksUsecase,
		updateBookUsecase:      updateBookUsecase,
		deleteBookUsecase:      deleteBookUsecase,
		searchBooksUsecase:     searchBooksUsecase,
		adjustStockUsecase:     adjustStockUsecase,
		changePriceUsecase:     changePriceUsecase,
		toggleActiveUsecase:    toggleActiveUsecase,
		validator:              validatorpkg.New(),
	}
}

// ============================================================================
// AUTHOR HANDLERS
// ============================================================================

// CreateAuthor godoc
// @Summary      Create a new author
// @Description  Create a new author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        request body application.CreateAuthorDTO true "Author creation data"
// @Success      201 {object} application.AuthorResponseDTO
// @Failure      400 {object} response.ErrorResponse "Validation failed"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/authors [post]
func (h *BookHandler) CreateAuthor(c *gin.Context) {
	var req application.CreateAuthorDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.createAuthorUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusCreated, result)
}

// GetAuthor godoc
// @Summary      Get author by ID
// @Description  Get author by ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id path string true "Author ID (UUID)"
// @Success      200 {object} application.AuthorResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Author not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/authors/{id} [get]
func (h *BookHandler) GetAuthor(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	result, err := h.getAuthorByIDUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// UpdateAuthor godoc
// @Summary      Update author
// @Description  Update author by ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id path string true "Author ID (UUID)"
// @Param        request body application.UpdateAuthorDTO true "Author update data"
// @Success      200 {object} application.AuthorResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid request"
// @Failure      404 {object} response.ErrorResponse "Author not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/authors/{id} [put]
func (h *BookHandler) UpdateAuthor(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	var req application.UpdateAuthorDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.updateAuthorUsecase.Execute(ctx, id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// DeleteAuthor godoc
// @Summary      Delete author
// @Description  Delete author by ID
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id path string true "Author ID (UUID)"
// @Success      204
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Author not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/authors/{id} [delete]
func (h *BookHandler) DeleteAuthor(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	err = h.deleteAuthorUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ============================================================================
// CATEGORY HANDLERS
// ============================================================================

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Create a new category
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        request body application.CreateCategoryDTO true "Category creation data"
// @Success      201 {object} application.CategoryResponseDTO
// @Failure      400 {object} response.ErrorResponse "Validation failed"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/categories [post]
func (h *BookHandler) CreateCategory(c *gin.Context) {
	var req application.CreateCategoryDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.createCategoryUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusCreated, result)
}

// GetCategories godoc
// @Summary      Get all categories
// @Description  Get all categories with pagination
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number (default: 1)" default(1)
// @Param        page_size query int false "Page size (default: 10)" default(10)
// @Success      200 {object} application.PaginatedResponse[application.CategoryResponseDTO]
// @Failure      400 {object} response.ErrorResponse "Invalid parameters"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/categories [get]
func (h *BookHandler) GetCategories(c *gin.Context) {
	var req application.GetCategoriesDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Set defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	ctx := c.Request.Context()
	result, err := h.getCategoriesUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// GetCategory godoc
// @Summary      Get category by ID
// @Description  Get category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "Category ID (UUID)"
// @Success      200 {object} application.CategoryResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Category not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/categories/{id} [get]
func (h *BookHandler) GetCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	result, err := h.getCategoryByIDUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// UpdateCategory godoc
// @Summary      Update category
// @Description  Update category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "Category ID (UUID)"
// @Param        request body application.UpdateCategoryDTO true "Category update data"
// @Success      200 {object} application.CategoryResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid request"
// @Failure      404 {object} response.ErrorResponse "Category not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/categories/{id} [put]
func (h *BookHandler) UpdateCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	var req application.UpdateCategoryDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.updateCategoryUsecase.Execute(ctx, id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// DeleteCategory godoc
// @Summary      Delete category
// @Description  Delete category by ID
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id path string true "Category ID (UUID)"
// @Success      204
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Category not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/categories/{id} [delete]
func (h *BookHandler) DeleteCategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	err = h.deleteCategoryUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ============================================================================
// PUBLISHER HANDLERS
// ============================================================================

// CreatePublisher godoc
// @Summary      Create a new publisher
// @Description  Create a new publisher
// @Tags         publishers
// @Accept       json
// @Produce      json
// @Param        request body application.CreatePublisherDTO true "Publisher creation data"
// @Success      201 {object} application.PublisherResponseDTO
// @Failure      400 {object} response.ErrorResponse "Validation failed"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/publishers [post]
func (h *BookHandler) CreatePublisher(c *gin.Context) {
	var req application.CreatePublisherDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.createPublisherUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusCreated, result)
}

// GetPublishers godoc
// @Summary      Get all publishers
// @Description  Get all publishers with pagination
// @Tags         publishers
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number (default: 1)" default(1)
// @Param        page_size query int false "Page size (default: 10)" default(10)
// @Success      200 {object} application.PaginatedResponse[application.PublisherResponseDTO]
// @Failure      400 {object} response.ErrorResponse "Invalid parameters"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/publishers [get]
func (h *BookHandler) GetPublishers(c *gin.Context) {
	var req application.GetPublishersDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Set defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	ctx := c.Request.Context()
	result, err := h.getPublishersUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// GetPublisher godoc
// @Summary      Get publisher by ID
// @Description  Get publisher by ID
// @Tags         publishers
// @Accept       json
// @Produce      json
// @Param        id path string true "Publisher ID (UUID)"
// @Success      200 {object} application.PublisherResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Publisher not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/publishers/{id} [get]
func (h *BookHandler) GetPublisher(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	result, err := h.getPublisherByIDUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// UpdatePublisher godoc
// @Summary      Update publisher
// @Description  Update publisher by ID
// @Tags         publishers
// @Accept       json
// @Produce      json
// @Param        id path string true "Publisher ID (UUID)"
// @Param        request body application.UpdatePublisherDTO true "Publisher update data"
// @Success      200 {object} application.PublisherResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid request"
// @Failure      404 {object} response.ErrorResponse "Publisher not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/publishers/{id} [put]
func (h *BookHandler) UpdatePublisher(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	var req application.UpdatePublisherDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.updatePublisherUsecase.Execute(ctx, id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// DeletePublisher godoc
// @Summary      Delete publisher
// @Description  Delete publisher by ID
// @Tags         publishers
// @Accept       json
// @Produce      json
// @Param        id path string true "Publisher ID (UUID)"
// @Success      204
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Publisher not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/publishers/{id} [delete]
func (h *BookHandler) DeletePublisher(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	err = h.deletePublisherUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ============================================================================
// BOOK HANDLERS
// ============================================================================

// CreateBook godoc
// @Summary      Create a new book
// @Description  Create a new book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        request body application.CreateBookDTO true "Book creation data"
// @Success      201 {object} application.BookResponseDTO
// @Failure      400 {object} response.ErrorResponse "Validation failed"
// @Failure      404 {object} response.ErrorResponse "Related entity not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req application.CreateBookDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.createBookUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusCreated, result)
}

// GetBook godoc
// @Summary      Get book by ID
// @Description  Get book by ID with author, category and publisher details
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID (UUID)"
// @Success      200 {object} application.BookResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Book not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	result, err := h.getBookByIDUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// GetBooks godoc
// @Summary      Get all books
// @Description  Get all books with optional filtering and pagination
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number (default: 1)" default(1)
// @Param        page_size query int false "Page size (default: 10)" default(10)
// @Param        category_id query string false "Filter by category ID"
// @Param        author_id query string false "Filter by author ID"
// @Param        active_only query boolean false "Get only active books"
// @Success      200 {object} application.PaginatedResponse[application.BookResponseDTO]
// @Failure      400 {object} response.ErrorResponse "Invalid parameters"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	var req application.GetBooksDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Set defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	ctx := c.Request.Context()
	result, err := h.getBooksUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// UpdateBook godoc
// @Summary      Update book
// @Description  Update book by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID (UUID)"
// @Param        request body application.UpdateBookDTO true "Book update data"
// @Success      200 {object} application.BookResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid request"
// @Failure      404 {object} response.ErrorResponse "Book not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	var req application.UpdateBookDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.updateBookUsecase.Execute(ctx, id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// DeleteBook godoc
// @Summary      Delete book
// @Description  Delete book by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID (UUID)"
// @Success      204
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Book not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	err = h.deleteBookUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// SearchBooks godoc
// @Summary      Search books
// @Description  Search books by title or description
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        query query string true "Search query"
// @Param        page query int false "Page number (default: 1)" default(1)
// @Param        page_size query int false "Page size (default: 10)" default(10)
// @Success      200 {object} application.PaginatedResponse[application.BookResponseDTO]
// @Failure      400 {object} response.ErrorResponse "Invalid parameters"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/search [get]
func (h *BookHandler) SearchBooks(c *gin.Context) {
	var req application.SearchBooksDTO
	if err := c.ShouldBindQuery(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err.Error())
		return
	}

	// Set defaults
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	ctx := c.Request.Context()
	result, err := h.searchBooksUsecase.Execute(ctx, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// AdjustStock godoc
// @Summary      Adjust book stock
// @Description  Adjust book stock (increase or decrease)
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID (UUID)"
// @Param        request body application.AdjustStockDTO true "Stock adjustment data"
// @Success      200 {object} application.BookResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid request"
// @Failure      404 {object} response.ErrorResponse "Book not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/{id}/adjust-stock [patch]
func (h *BookHandler) AdjustStock(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	var req application.AdjustStockDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.adjustStockUsecase.Execute(ctx, id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// ChangePrice godoc
// @Summary      Change book price
// @Description  Change book price
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID (UUID)"
// @Param        request body application.ChangePriceDTO true "New price data"
// @Success      200 {object} application.BookResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid request"
// @Failure      404 {object} response.ErrorResponse "Book not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/{id}/change-price [patch]
func (h *BookHandler) ChangePrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	var req application.ChangePriceDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		validationErrors := validator.GetErrorValidate(err)
		response.WriteErrorResponse(c, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	ctx := c.Request.Context()
	result, err := h.changePriceUsecase.Execute(ctx, id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// ToggleActive godoc
// @Summary      Toggle book active status
// @Description  Toggle book active/inactive status
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID (UUID)"
// @Success      200 {object} application.BookResponseDTO
// @Failure      400 {object} response.ErrorResponse "Invalid ID"
// @Failure      404 {object} response.ErrorResponse "Book not found"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /books/{id}/toggle-active [patch]
func (h *BookHandler) ToggleActive(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.WriteErrorResponse(c, http.StatusBadRequest, "Invalid ID format", nil)
		return
	}

	ctx := c.Request.Context()
	result, err := h.toggleActiveUsecase.Execute(ctx, id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	response.WriteSuccessResponse(c, http.StatusOK, result)
}

// handleError xử lý lỗi từ domain/application layer
func (h *BookHandler) handleError(c *gin.Context, err error) {
	// Kiểm tra custom domain errors
	switch err {
	case domain.ErrBookNotFound:
		response.WriteErrorResponse(c, http.StatusNotFound, "Book not found", nil)
	case domain.ErrAuthorNotFound:
		response.WriteErrorResponse(c, http.StatusNotFound, "Author not found", nil)
	case domain.ErrCategoryNotFound:
		response.WriteErrorResponse(c, http.StatusNotFound, "Category not found", nil)
	case domain.ErrPublisherNotFound:
		response.WriteErrorResponse(c, http.StatusNotFound, "Publisher not found", nil)
	case domain.ErrInsufficientStock:
		response.WriteErrorResponse(c, http.StatusBadRequest, "Insufficient stock", nil)
	case domain.ErrBookTitleRequired:
		response.WriteErrorResponse(c, http.StatusBadRequest, domain.ErrBookTitleRequired.Error(), nil)
	case domain.ErrAuthorNameRequired:
		response.WriteErrorResponse(c, http.StatusBadRequest, domain.ErrAuthorNameRequired.Error(), nil)
	case domain.ErrCategoryNameRequired:
		response.WriteErrorResponse(c, http.StatusBadRequest, domain.ErrCategoryNameRequired.Error(), nil)
	case domain.ErrPublisherNameRequired:
		response.WriteErrorResponse(c, http.StatusBadRequest, domain.ErrPublisherNameRequired.Error(), nil)
	case domain.ErrCategoryNameExists:
		response.WriteErrorResponse(c, http.StatusConflict, domain.ErrCategoryNameExists.Error(), nil)
	case domain.ErrBookPriceInvalid:
		response.WriteErrorResponse(c, http.StatusBadRequest, domain.ErrBookPriceInvalid.Error(), nil)
	case domain.ErrBookStockInvalid:
		response.WriteErrorResponse(c, http.StatusBadRequest, domain.ErrBookStockInvalid.Error(), nil)
	default:
		// Kiểm tra wrapped errors
		if wrappedErr, ok := err.(*errs.AppError); ok {
			response.WriteErrorResponse(c, http.StatusInternalServerError, wrappedErr.Message, wrappedErr.Code)
		} else {
			response.WriteErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
		}
	}
}
