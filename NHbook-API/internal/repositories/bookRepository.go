package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book *models.Book, tx *gorm.DB) error
	GetBookByID(bookID uint) (*models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

// GetBookByID implements IBookRepository.
func (b *bookRepository) GetBookByID(bookID uint) (*models.Book, error) {
	panic("unimplemented")
}

// CreateBook implements IBookRepository.
func (b *bookRepository) CreateBook(book *models.Book, tx *gorm.DB) error {
	err := tx.Create(book).Error
	return err
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{
		db: db,
	}
}
