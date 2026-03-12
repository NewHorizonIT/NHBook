package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type GetAuthorByIDUsecase struct {
	repo domain.IAuthorRepository
}

func NewGetAuthorByIDUsecase(repo domain.IAuthorRepository) *GetAuthorByIDUsecase {
	return &GetAuthorByIDUsecase{repo: repo}
}

func (uc *GetAuthorByIDUsecase) Execute(ctx context.Context, id uuid.UUID) (*application.AuthorResponseDTO, error) {
	// 1. Get author from repository
	author, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, domain.ErrAuthorNotFound
	}

	// 2. Return response
	return application.ToAuthorResponse(author), nil
}
