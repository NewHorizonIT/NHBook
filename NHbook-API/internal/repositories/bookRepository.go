package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book *models.Book, tx *gorm.DB) error
	GetBookByID(bookID uint) (*models.Book, error)
	GetListBook(limit int, page int, categoryID int, authorID int) ([]models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

// GetListBook implements IBookRepository.
func (b *bookRepository) GetListBook(limit int, page int, categoryID int, authorID int) ([]models.Book, error) {
	var books []models.Book
	query := b.db

	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if authorID != 0 {
		query = query.Joins("JOIN book_author ON books.id = book_author.book_id").
			Where("book_author.author_id = ?", authorID).
			Distinct("books.*")

	}
	if err := query.Limit(limit).Offset((page - 1) * limit).Preload("Category").Preload("Authors").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
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
