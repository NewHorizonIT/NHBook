package services

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"gorm.io/gorm"
)

type IBookService interface {
	CreateBook(book *models.Book, tx *gorm.DB) (*models.Book, error)
	GetListBook(limit int, page int, categoryID int, authorID int) ([]models.Book, error)
}

type bookService struct {
	bookRepo repositories.IBookRepository
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
