package repositories

import (
	"errors"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(book *models.Book, tx *gorm.DB) error
	GetBookByID(bookID int) (*models.Book, error)
	GetListBook(query *request.QueryLimit, categoryID int, authorID int) ([]models.Book, error)
	IsExistBook(bookID int) (bool, error)
	GetStock(bookID int) (int, error)
	UpdateStock(tx *gorm.DB, bookID int, stock int) error
	GetTitleBookByID(tx *gorm.DB, bookID uint) (string, error)
	GetListBookByCategory(category int, query *request.QueryLimit) ([]models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

// GetListBookByCategory implements IBookRepository.
func (b *bookRepository) GetListBookByCategory(category int, query *request.QueryLimit) ([]models.Book, error) {
	var books []models.Book
	err := b.db.Model(&models.Book{}).Joins("categories").Where("categories.id = ?", category).Find(&books).Limit(query.Limit).Offset((query.Page - 1) * 20).Error

	if err != nil {
		return nil, err
	}

	return books, nil
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
func (b *bookRepository) GetListBook(query *request.QueryLimit, categoryID int, authorID int) ([]models.Book, error) {
	var books []models.Book
	queryDB := b.db.Model(&models.Book{})

	if categoryID != 0 {
		queryDB = queryDB.Where("category_id = ?", categoryID)
	}

	if authorID != 0 {
		queryDB = queryDB.Joins("JOIN book_author ON books.id = book_author.book_id").
			Where("book_author.author_id = ?", authorID).
			Distinct("books.*")
	}
	if err := queryDB.Limit(query.Limit).Offset((query.Page - 1) * query.Limit).Preload("Category").Preload("Authors").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// GetBookByID implements IBookRepository.
func (b *bookRepository) GetBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	if err := b.db.Where("id = ?", bookID).Preload("Category").Preload("Authors").First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil

}

// CreateBook implements IBookRepository.
func (b *bookRepository) CreateBook(book *models.Book, tx *gorm.DB) error {
	err := tx.Create(&book).Error
	return err
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{
		db: db,
	}
}
