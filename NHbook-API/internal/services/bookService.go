package services

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"gorm.io/gorm"
)

type IBookService interface {
	CreateBook(book *models.Book, tx *gorm.DB) (*response.CreateBookResponse, error)
}

type bookService struct {
	bookRepo repositories.IBookRepository
}

// CreateBook implements IBookService.
func (b *bookService) CreateBook(book *models.Book, tx *gorm.DB) (*response.CreateBookResponse, error) {
	err := b.bookRepo.CreateBook(book, tx)

	if err != nil {
		return nil, err
	}

	// Create response

	res := &response.CreateBookResponse{
		ID:          book.ID,
		Title:       book.Title,
		ImageURL:    book.ImageURL,
		Price:       book.Price,
		Description: book.Description,
		Stock:       book.Stock,
		CategoryID:  book.CategoryID,
		CreatedAt:   book.CreatedAt,
	}

	return res, nil

}

func NewBookService(bookRepo repositories.IBookRepository) IBookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}
