package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type AdjustStockUsecase struct {
	repo domain.IBookRepository
}

func NewAdjustStockUsecase(repo domain.IBookRepository) *AdjustStockUsecase {
	return &AdjustStockUsecase{repo: repo}
}

func (uc *AdjustStockUsecase) Execute(ctx context.Context, id uuid.UUID, dto *application.AdjustStockDTO) (*application.BookResponseDTO, error) {
	// 1. Check if book exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ErrBookNotFound
	}

	// 2. Adjust stock in repository
	book, err := uc.repo.AdjustStock(ctx, id, dto.Delta)
	if err != nil {
		return nil, err
	}

	// 3. Return response
	return application.ToBookResponse(book), nil
}
