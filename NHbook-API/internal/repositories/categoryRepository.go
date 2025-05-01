package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetAllcategory(status int) ([]models.Category, error)
	GetCategoryByID(categoryID uint) (*models.Category, error)
	CheckCategoryIsExists(categoryID int, tx *gorm.DB) (*models.Category, error)
	CategoryIsExitsByName(categoryName string) (bool, error)
	Updatecategory(category *models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

// Updatecategory implements ICategoryRepository.
func (c *categoryRepository) Updatecategory(category *models.Category) error {
	if err := c.db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(&category).First(&category).Error; err != nil {
		return err
	}

	return nil
}

// CategoryIsExits implements ICategoryRepository.
func (c *categoryRepository) CategoryIsExitsByName(categoryName string) (bool, error) {
	if err := c.db.Model(&models.Category{}).Where("name = ?", categoryName).Error; err != nil {
		return false, err
	}

	return true, nil
}

// GetAllcategory implements ICategoryRepository.
func (c *categoryRepository) GetAllcategory(status int) ([]models.Category, error) {
	var categories []models.Category
	switch status {
	case 0, 1:
		if err := c.db.Where("status = ?", status).Find(&categories).Error; err != nil {
			return nil, err
		}
	default:
		if err := c.db.Find(&categories).Error; err != nil {
			return nil, err
		}
	}

	return categories, nil
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
	return c.db.Create(&category).Error
}

// GetCategoryByNameByID implements ICategoryRepository.
func (c *categoryRepository) GetCategoryByID(categoryID uint) (*models.Category, error) {
	var category models.Category

	err := c.db.Model(&models.Category{}).Where("id = ?", categoryID).First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
