package services

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models/common/response"
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/repositories"
	"gorm.io/gorm"
)

type IAuthorService interface {
	CreateAuthor(author *models.Author) (*response.AuthorResponse, error)
	CheckAuthorExists(authorName string, tx *gorm.DB) (*models.Author, error)
}

type authorService struct {
	authorRepo repositories.IAuthorRepository
}

// CheckAuthorExits implements IAuthorService.
func (a *authorService) CheckAuthorExists(authorName string, tx *gorm.DB) (*models.Author, error) {

	author, err := a.authorRepo.GetOrCreateAuthor(authorName, tx)

	if err != nil {
		return nil, err
	}

	return author, nil
}

// CreateAuthor implements IAuthorService.
func (a *authorService) CreateAuthor(author *models.Author) (*response.AuthorResponse, error) {
	panic("unimplemented")
}

func NewAuthorService(authorRepo repositories.IAuthorRepository) IAuthorService {
	return &authorService{
		authorRepo: authorRepo,
	}
}
