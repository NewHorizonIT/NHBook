package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
)

type CreateAuthorUsecase struct {
	repo domain.IAuthorRepository
}

func NewCreateAuthorUsecase(repo domain.IAuthorRepository) *CreateAuthorUsecase {
	return &CreateAuthorUsecase{repo: repo}
}

func (uc *CreateAuthorUsecase) Execute(ctx context.Context, dto *application.CreateAuthorDTO) (*application.AuthorResponseDTO, error) {
	// 1. Create domain entity
	author, err := domain.NewAuthor(dto.Name, dto.Bio)
	if err != nil {
		return nil, err
	}

	// 2. Save to repository
	if err := uc.repo.Create(ctx, author); err != nil {
		return nil, err
	}

	// 3. Return response
	return application.ToAuthorResponse(author), nil
}
