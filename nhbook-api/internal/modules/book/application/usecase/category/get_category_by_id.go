package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type GetCategoryByIDUsecase struct {
	repo domain.ICategoryRepository
}

func NewGetCategoryByIDUsecase(repo domain.ICategoryRepository) *GetCategoryByIDUsecase {
	return &GetCategoryByIDUsecase{repo: repo}
}

func (uc *GetCategoryByIDUsecase) Execute(ctx context.Context, id uuid.UUID) (*application.CategoryResponseDTO, error) {
	// 1. Get category from repository
	category, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, domain.ErrCategoryNotFound
	}

	// 2. Return response
	return application.ToCategoryResponse(category), nil
}
