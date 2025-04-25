package services

import (
	"errors"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"gorm.io/gorm"
)

var (
	ErrBookExist = errors.New("book isn't exist")
)

type IBookService interface {
	CreateBook(book *models.Book, tx *gorm.DB) (*models.Book, error)
	GetListBookByCategory(category int, query *request.QueryLimit) (*response.GetBookResponse, error)
	GetListBook(limit int, page int, categoryID int, authorID int) ([]models.Book, error)
	CheckBookExist(bookID int) (bool, error)
	GetBookByID(bookID uint) (*models.Book, error)
}

type bookService struct {
	bookRepo repositories.IBookRepository
}

func (b *bookService) GetListBookByCategory(category int, query *request.QueryLimit) (*response.GetBookResponse, error) {
	books, err := b.bookRepo.GetListBookByCategory(category, query)

	if err != nil {
		return nil, err
	}

	// Create GetListBookByCategoryResponse
	res := &response.GetBookResponse{
		Limit: query.Limit,
		Page:  query.Page,
		Data:  books,
	}
	return res, nil
}

// GetBookByID implements IBookService.
func (b *bookService) GetBookByID(bookID uint) (*models.Book, error) {
	book, err := b.bookRepo.GetBookByID(bookID)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, ErrBookExist
	}

	return book, nil
}

// CheckBookExist implements IBookService.
func (b *bookService) CheckBookExist(bookID int) (bool, error) {
	isExist, err := b.bookRepo.IsExistBook(bookID)

	if err != nil {
		return false, err
	}

	if !isExist {
		return false, ErrBookExist
	}

	return isExist, nil
}

// GetListBook implements IBookService.
func (b *bookService) GetListBook(limit int, page int, categoryID int, authorID int) ([]models.Book, error) {
	var books []models.Book

	books, err := b.bookRepo.GetListBook(limit, page, categoryID, authorID)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// CreateBook implements IBookService.
func (b *bookService) CreateBook(book *models.Book, tx *gorm.DB) (*models.Book, error) {
	err := b.bookRepo.CreateBook(book, tx)
	if err != nil {
		return nil, err
	}
	// Create response
	res := book

	return res, nil

}

func NewBookService(bookRepo repositories.IBookRepository) IBookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}
