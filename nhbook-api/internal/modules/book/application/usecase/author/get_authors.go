package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
)

type GetAuthorsUsecase struct {
	repo domain.IAuthorRepository
}

func NewGetAuthorsUsecase(repo domain.IAuthorRepository) *GetAuthorsUsecase {
	return &GetAuthorsUsecase{repo: repo}
}

func (uc *GetAuthorsUsecase) Execute(ctx context.Context, dto *application.GetAuthorsDTO) (*application.PaginatedResponse[application.AuthorResponseDTO], error) {
	// 1. Create pagination
	pagination := shared.NewPagination(dto.Page, dto.PageSize)

	// 2. Get authors from repository
	authors, err := uc.repo.GetAll(ctx, pagination)
	if err != nil {
		return nil, err
	}

	// 3. Get total count
	totalCount, err := uc.repo.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	// 4. Create paginated result
	result := shared.NewPaginatedResult(authors, totalCount, pagination)

	// 5. Map to response
	response := application.ToPaginatedResponse(result, application.ToAuthorResponseList)
	return &response, nil
}
