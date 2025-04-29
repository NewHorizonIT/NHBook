package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategoryByNameByID(categoryID uint) (*models.Category, error)
	GetCategoryByName(categoryName string) (*models.Category, error)
	CheckCategoryIsExists(categoryID int, tx *gorm.DB) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

// CheckCategoryIsExists implements ICategoryRepository.
func (c *categoryRepository) CheckCategoryIsExists(categoryID int, tx *gorm.DB) (*models.Category, error) {
	var category models.Category
	err := tx.Where("id = ?", categoryID).First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
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
