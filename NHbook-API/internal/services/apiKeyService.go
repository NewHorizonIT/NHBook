package services

import "github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"

type IApiKeyService interface {
	CheckApiKey(key string) (bool, error)
}

type apiKeyService struct {
	ApiKeyRepo repositories.IApiKeyRepository
}

func NewApiKeyService(repo repositories.IApiKeyRepository) IApiKeyService {
	return &apiKeyService{
		ApiKeyRepo: repo,
	}
}

// CheckApiKey implements IApiKeyService.
func (as *apiKeyService) CheckApiKey(key string) (bool, error) {
	isExits, err := as.ApiKeyRepo.IsExitsApiKey(key)
	if err != nil {
		return false, err
	}
	if isExits {
		return true, nil
	}
	return false, nil
}
