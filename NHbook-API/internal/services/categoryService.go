package services

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"gorm.io/gorm"
)

type ICategoryService interface {
	CreateCategory(category *models.Category) (*response.CategoryResponse, error)
	CheckCategoryExitsByID(categoryID int, tx *gorm.DB) (*models.Category, error)
	GetCategoryIDByName(category string) (int, error)
}

type categoryService struct {
	categoryRepo repositories.ICategoryRepository
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
func (c *categoryService) CreateCategory(category *models.Category) (*response.CategoryResponse, error) {
	panic("unimplemented")
}

// GetCategoryIDByName implements ICategoryService.
func (c *categoryService) GetCategoryIDByName(category string) (int, error) {
	CategoryFound, err := c.categoryRepo.GetCategoryByName(category)

	if err != nil {
		return 0, err
	}

	return int(CategoryFound.ID), nil
}

func NewCategoryService(categoryRepo repositories.ICategoryRepository) ICategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}
