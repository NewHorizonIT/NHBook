package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type ToggleActiveUsecase struct {
	repo domain.IBookRepository
}

func NewToggleActiveUsecase(repo domain.IBookRepository) *ToggleActiveUsecase {
	return &ToggleActiveUsecase{repo: repo}
}

func (uc *ToggleActiveUsecase) Execute(ctx context.Context, id uuid.UUID) (*application.BookResponseDTO, error) {
	// 1. Check if book exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ErrBookNotFound
	}

	// 2. Toggle active status in repository
	book, err := uc.repo.ToggleActive(ctx, id)
	if err != nil {
		return nil, err
	}

	// 3. Return response
	return application.ToBookResponse(book), nil
}
