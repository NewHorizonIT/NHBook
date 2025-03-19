package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type ITokenRepository interface {
	CreateOrUpdateToken(userID string, token string) error
	GetTokenByID(userID string) (models.Token, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) ITokenRepository {
	return &tokenRepository{
		db: db,
	}
}

// CreateOrUpdatedToken implements ITokenRepository.
func (t *tokenRepository) CreateOrUpdateToken(userID string, token string) error {
	newToken := models.Token{
		UserID: userID,
		Token:  token,
	}
	return t.db.Save(&newToken).Error
}

// GetTokenByID implements ITokenRepository.
func (t *tokenRepository) GetTokenByID(userID string) (models.Token, error) {
	var foundToken models.Token
	err := t.db.Model(&models.Token{}).Where("user_id = ?", userID).First(&foundToken).Error
	return foundToken, err
}
