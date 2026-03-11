package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
)

type CreateCategoryUsecase struct {
	repo domain.ICategoryRepository
}

func NewCreateCategoryUsecase(repo domain.ICategoryRepository) *CreateCategoryUsecase {
	return &CreateCategoryUsecase{repo: repo}
}

func (uc *CreateCategoryUsecase) Execute(ctx context.Context, dto *application.CreateCategoryDTO) (*application.CategoryResponseDTO, error) {
	// 1. Check if category name already exists
	exists, err := uc.repo.ExistsByName(ctx, dto.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrCategoryNameExists
	}

	// 2. Create domain entity
	category, err := domain.NewCategory(dto.Name)
	if err != nil {
		return nil, err
	}

	// 3. Save to repository
	if err := uc.repo.Create(ctx, category); err != nil {
		return nil, err
	}

	// 4. Return response
	return application.ToCategoryResponse(category), nil
}
