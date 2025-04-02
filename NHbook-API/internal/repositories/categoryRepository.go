package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategoryByNameByID(categoryID uint) (*models.Category, error)
	GetCategoryByName(categoryName string) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements ICategoryRepository.
func (c *categoryRepository) CreateCategory(category *models.Category) error {
	panic("unimplemented")
}

// GetCategoryByName implements ICategoryRepository.
func (c *categoryRepository) GetCategoryByName(categoryName string) (*models.Category, error) {
	panic("unimplemented")
}

// GetCategoryByNameByID implements ICategoryRepository.
func (c *categoryRepository) GetCategoryByNameByID(categoryID uint) (*models.Category, error) {
	panic("unimplemented")
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
