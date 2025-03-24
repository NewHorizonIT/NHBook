package repositories

import (
	"errors"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IApiKeyRepository interface {
	GetApiKeyByKey(key string) (*models.ApiKey, error)
	IsExitsApiKey(key string) (bool, error)
}

type ApiKeyRepository struct {
	db *gorm.DB
}

func NewApiKeyRepository(db *gorm.DB) IApiKeyRepository {
	return &ApiKeyRepository{
		db: db,
	}
}

// GetApiKeyByKey implements IApiKeyRepository.
func (a *ApiKeyRepository) GetApiKeyByKey(key string) (*models.ApiKey, error) {
	var apiKey *models.ApiKey
	if err := a.db.Where("`apikey = ?", key).First(apiKey).Error; err != nil {
		return nil, err
	}
	return apiKey, nil

}

// IsExitsApiKey implements IApiKeyRepository.
func (a *ApiKeyRepository) IsExitsApiKey(key string) (bool, error) {
	var apiKey models.ApiKey
	err := a.db.Where("api_key = ?", key).First(&apiKey).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
