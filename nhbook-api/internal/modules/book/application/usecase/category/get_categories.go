package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
)

type GetCategoriesUsecase struct {
	repo domain.ICategoryRepository
}

func NewGetCategoriesUsecase(repo domain.ICategoryRepository) *GetCategoriesUsecase {
	return &GetCategoriesUsecase{repo: repo}
}

func (uc *GetCategoriesUsecase) Execute(ctx context.Context, dto *application.GetCategoriesDTO) (*application.PaginatedResponse[application.CategoryResponseDTO], error) {
	// 1. Create pagination
	pagination := shared.NewPagination(dto.Page, dto.PageSize)

	// 2. Get categories from repository
	categories, err := uc.repo.GetAll(ctx, pagination)
	if err != nil {
		return nil, err
	}

	// 3. Get total count
	totalCount, err := uc.repo.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	// 4. Create paginated result
	result := shared.NewPaginatedResult(categories, totalCount, pagination)

	// 5. Map to response
	response := application.ToPaginatedResponse(result, application.ToCategoryResponseList)
	return &response, nil
}
