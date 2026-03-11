package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type UpdateAuthorUsecase struct {
	repo domain.IAuthorRepository
}

func NewUpdateAuthorUsecase(repo domain.IAuthorRepository) *UpdateAuthorUsecase {
	return &UpdateAuthorUsecase{repo: repo}
}

func (uc *UpdateAuthorUsecase) Execute(ctx context.Context, id uuid.UUID, dto *application.UpdateAuthorDTO) (*application.AuthorResponseDTO, error) {
	// 1. Get existing author
	author, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, domain.ErrAuthorNotFound
	}

	// 2. Update domain entity
	author.Update(dto.Name, dto.Bio)

	// 3. Save to repository
	if err := uc.repo.Update(ctx, author); err != nil {
		return nil, err
	}

	// 4. Return response
	return application.ToAuthorResponse(author), nil
}
