package services

import (
	"strings"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/request"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ICategoryService interface {
	CreateCategory(category *request.CategoryRequest) (*models.Category, error)
	CheckCategoryExitsByID(categoryID int, tx *gorm.DB) (*models.Category, error)
	GetCategoryIDByName(category string) (int, error)
	GetCategoryByID(categoryID int) (*models.Category, error)
	GetAllCategory(status int) ([]models.Category, error)
	UpdateCategory(category *request.CategoryUpdate) (*models.Category, error)
}

type categoryService struct {
	categoryRepo repositories.ICategoryRepository
}

// UpdateCategory implements ICategoryService.
func (c *categoryService) UpdateCategory(category *request.CategoryUpdate) (*models.Category, error) {
	categoryUpdate := &models.Category{}
	copier.Copy(&categoryUpdate, &category)
	if err := c.categoryRepo.Updatecategory(categoryUpdate); err != nil {
		return nil, err
	}

	return categoryUpdate, nil
}

// GetAllCategory implements ICategoryService.
func (c *categoryService) GetAllCategory(status int) ([]models.Category, error) {
	categories, err := c.categoryRepo.GetAllcategory(status)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByID implements ICategoryService.
func (c *categoryService) GetCategoryByID(categoryID int) (*models.Category, error) {
	category, err := c.categoryRepo.GetCategoryByID(uint(categoryID))

	if err != nil {
		return nil, err
	}

	return category, nil
}

// CheckCategoryExitsByID implements ICategoryService.
func (c *categoryService) CheckCategoryExitsByID(categoryID int, tx *gorm.DB) (*models.Category, error) {
	category, err := c.categoryRepo.CheckCategoryIsExists(categoryID, tx)

	if err != nil {
		return nil, err
	}

	return category, nil
}

// CreateCategory implements ICategoryService.
func (c *categoryService) CreateCategory(category *request.CategoryRequest) (*models.Category, error) {
	// Step 1: Check Category is Exists
	category.Name = strings.ToLower(strings.TrimSpace(category.Name))
	categoryFound, err := c.categoryRepo.CategoryIsExitsByName(category.Name)

	if err != nil || !categoryFound {
		return nil, err
	}

	// Step 3: create model category
	newCategory := &models.Category{}
	copier.Copy(newCategory, category)

	// Step 2: Create category
	if err := c.categoryRepo.CreateCategory(newCategory); err != nil {
		return nil, err
	}

	return newCategory, nil

}

// GetCategoryIDByName implements ICategoryService.
func (c *categoryService) GetCategoryIDByName(category string) (int, error) {
	return 0, nil
}

func NewCategoryService(categoryRepo repositories.ICategoryRepository) ICategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}
