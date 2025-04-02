package repositories

import (
	"time"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IAuthorRepository interface {
	CreateAuthor(author *models.Author) error
	GetAuthorByName(authorName string) (*models.Author, error)
	GetAuthorByID(authorID uint) (*models.Author, error)
	GetOrCreateAuthor(authorName string, tx *gorm.DB) (*models.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

// GetAuthorByID implements IAuthorRepository.
func (a *authorRepository) GetAuthorByID(authorID uint) (*models.Author, error) {
	panic("unimplemented")
}

// GetOrCreateAuthorID implements IAuthorRepository.
func (a *authorRepository) GetOrCreateAuthor(authorName string, tx *gorm.DB) (*models.Author, error) {
	if tx == nil {
		tx = a.db
	}

	var author models.Author
	err := tx.Where("name = ?", authorName).FirstOrCreate(&author, models.Author{
		Name:      authorName,
		Bio:       "",
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}).Error

	if err != nil {
		return nil, err
	}
	return &author, nil
}

// CreateAuthor implements IAuthorRepository.
func (a *authorRepository) CreateAuthor(author *models.Author) error {
	panic("unimplemented")
}

// GetAuthorByName implements IAuthorRepository.
func (a *authorRepository) GetAuthorByName(authorName string) (*models.Author, error) {
	panic("unimplemented")
}

func NewAuthorRepository(db *gorm.DB) IAuthorRepository {
	return &authorRepository{
		db: db,
	}
}
