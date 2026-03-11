package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
)

type GetPublishersUsecase struct {
	repo domain.IPublisherRepository
}

func NewGetPublishersUsecase(repo domain.IPublisherRepository) *GetPublishersUsecase {
	return &GetPublishersUsecase{repo: repo}
}

func (uc *GetPublishersUsecase) Execute(ctx context.Context, dto *application.GetPublishersDTO) (*application.PaginatedResponse[application.PublisherResponseDTO], error) {
	// 1. Create pagination
	pagination := shared.NewPagination(dto.Page, dto.PageSize)

	// 2. Get publishers from repository
	publishers, err := uc.repo.GetAll(ctx, pagination)
	if err != nil {
		return nil, err
	}

	// 3. Get total count
	totalCount, err := uc.repo.GetCount(ctx)
	if err != nil {
		return nil, err
	}

	// 4. Create paginated result
	result := shared.NewPaginatedResult(publishers, totalCount, pagination)

	// 5. Map to response
	response := application.ToPaginatedResponse(result, application.ToPublisherResponseList)
	return &response, nil
}
