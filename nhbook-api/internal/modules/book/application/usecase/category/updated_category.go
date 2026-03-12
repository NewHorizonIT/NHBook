package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type UpdateCategoryUsecase struct {
	repo domain.ICategoryRepository
}

func NewUpdateCategoryUsecase(repo domain.ICategoryRepository) *UpdateCategoryUsecase {
	return &UpdateCategoryUsecase{repo: repo}
}

func (uc *UpdateCategoryUsecase) Execute(ctx context.Context, id uuid.UUID, dto *application.UpdateCategoryDTO) (*application.CategoryResponseDTO, error) {
	// 1. Get existing category
	category, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, domain.ErrCategoryNotFound
	}

	// 2. Check if new name already exists (if name is being changed)
	if dto.Name != category.Name {
		exists, err := uc.repo.ExistsByName(ctx, dto.Name)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, domain.ErrCategoryNameExists
		}
	}

	// 3. Update domain entity
	if err := category.Update(dto.Name); err != nil {
		return nil, err
	}

	// 4. Save to repository
	if err := uc.repo.Update(ctx, category); err != nil {
		return nil, err
	}

	// 5. Return response
	return application.ToCategoryResponse(category), nil
}
