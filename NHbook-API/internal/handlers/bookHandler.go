package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/services"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/utils"
	"github.com/gin-gonic/gin"
)

var (
	ErrBindBody    = errors.New("binding body error")
	ErrTransaction = errors.New("transaction error")
)

type BookHandler struct {
	bookService     services.IBookService
	authorService   services.IAuthorService
	categoryService services.ICategoryService
}

func NewBookHandler(bs services.IBookService, as services.IAuthorService, cs services.ICategoryService) *BookHandler {
	return &BookHandler{
		bookService:     bs,
		authorService:   as,
		categoryService: cs,
	}
}

func (bh *BookHandler) CreateBook(c *gin.Context) {
	var req request.CreateBookRequest
	// Step 1: Binding body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrBindBody, err).Error())
		return
	}

	// Step 2: Create Book
	tx := global.MySQL.Begin()

	// Rollback if any error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			utils.WriteError(c, http.StatusInternalServerError, "Unexpected error")
		}
	}()
	// Get  author
	authors := req.Authors
	var listAuthor []models.Author

	for _, author := range authors {
		a, err := bh.authorService.CheckAuthorExists(author, tx)
		if err != nil {
			utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrTransaction, err).Error())
			return
		}
		listAuthor = append(listAuthor, *a)
	}

	// Check Category exist
	categoryID := req.CategoryID

	name, err := bh.categoryService.CheckCategoryExitsByID(categoryID, tx)

	if err != nil || name == "" {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrTransaction, err).Error())
		return
	}

	publishedAt, err := time.Parse(time.DateOnly, req.PublishedAt)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}
	// Create new Book
	newBook := models.Book{
		Title:       req.Title,
		Authors:     listAuthor,
		ImageURL:    req.ImageURL,
		Price:       req.Price,
		Description: req.Description,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		PublishedAt: publishedAt,
	}

	book, err := bh.bookService.CreateBook(&newBook, tx)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, utils.FormatError(ErrTransaction, err).Error())
		tx.Rollback()
		return
	}

	metadata := &response.CreateBookResponse{
		ID:          book.ID,
		Title:       book.Title,
		ImageURL:    book.ImageURL,
		Price:       book.Price,
		Description: book.Description,
		PublishedAt: book.PublishedAt.Format(time.DateOnly),
		Authors:     req.Authors,
		Stock:       book.Stock,
		Category:    name,
		CategoryID:  book.CategoryID,
		CreatedAt:   book.CreatedAt,
	}

	tx.Commit()
	// Step 2: Create book
	utils.WriteResponse(c, http.StatusOK, "Create book Success", metadata, nil)

}
