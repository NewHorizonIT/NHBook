package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type GetBookByIDUsecase struct {
	repo domain.IBookRepository
}

func NewGetBookByIDUsecase(repo domain.IBookRepository) *GetBookByIDUsecase {
	return &GetBookByIDUsecase{repo: repo}
}

func (uc *GetBookByIDUsecase) Execute(ctx context.Context, id uuid.UUID) (*application.BookResponseDTO, error) {
	// 1. Get book with relations from repository
	book, err := uc.repo.GetByIDWithRelations(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, domain.ErrBookNotFound
	}

	// 2. Return response
	return application.ToBookWithRelationsResponse(book), nil
}
