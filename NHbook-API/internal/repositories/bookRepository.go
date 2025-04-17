package repositories

import (
	"errors"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book *models.Book, tx *gorm.DB) error
	GetBookByID(bookID uint) (*models.Book, error)
	GetListBook(limit int, page int, categoryID int, authorID int) ([]models.Book, error)
	IsExistBook(bookID int) (bool, error)
	GetStock(bookID int) (int, error)
	UpdateStock(tx *gorm.DB, bookID int, stock int) error
	GetTitleBookByID(tx *gorm.DB, bookID uint) (string, error)
}

type bookRepository struct {
	db *gorm.DB
}

// GetTitleBookByID implements IBookRepository.
func (b *bookRepository) GetTitleBookByID(tx *gorm.DB, bookID uint) (string, error) {
	var book models.Book
	if err := tx.Model(&models.Book{}).Where("id = ?", bookID).Select("title").Scan(&book).Error; err != nil {
		return "", err
	}

	return book.Title, nil
}

// UpdateStock implements IBookRepository.
func (b *bookRepository) UpdateStock(tx *gorm.DB, bookID int, stock int) error {
	var currentStock int
	if err := tx.Model(&models.Book{}).Where("id = ?", bookID).Select("stock").Scan(&currentStock).Error; err != nil {
		return err
	}
	newStock := currentStock - stock
	return tx.Model(&models.Book{}).Where("id = ?", bookID).Update("stock", newStock).Error
}

// GetStock implements IBookRepository.
func (b *bookRepository) GetStock(bookID int) (int, error) {
	var stock int
	err := b.db.Model(models.Book{}).Where("id = ?", bookID).First(&models.Book{}).Select("stock").Scan(&stock).Error

	if err != nil {
		return -1, err
	}

	return stock, nil
}

// FindBookByID implements IBookRepository.
func (b *bookRepository) IsExistBook(bookID int) (bool, error) {
	err := b.db.Where("id = ?", bookID).First(&models.Book{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
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
