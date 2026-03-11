package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type ChangePriceUsecase struct {
	repo domain.IBookRepository
}

func NewChangePriceUsecase(repo domain.IBookRepository) *ChangePriceUsecase {
	return &ChangePriceUsecase{repo: repo}
}

func (uc *ChangePriceUsecase) Execute(ctx context.Context, id uuid.UUID, dto *application.ChangePriceDTO) (*application.BookResponseDTO, error) {
	// 1. Check if book exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ErrBookNotFound
	}

	// 2. Change price in repository
	book, err := uc.repo.ChangePrice(ctx, id, dto.Price)
	if err != nil {
		return nil, err
	}

	// 3. Return response
	return application.ToBookResponse(book), nil
}
