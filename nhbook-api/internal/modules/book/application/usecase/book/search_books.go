package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
)

type SearchBooksUsecase struct {
	repo domain.IBookRepository
}

func NewSearchBooksUsecase(repo domain.IBookRepository) *SearchBooksUsecase {
	return &SearchBooksUsecase{repo: repo}
}

func (uc *SearchBooksUsecase) Execute(ctx context.Context, dto *application.SearchBooksDTO) (*application.PaginatedResponse[application.BookResponseDTO], error) {
	// 1. Create pagination
	pagination := shared.NewPagination(dto.Page, dto.PageSize)

	// 2. Search books
	books, err := uc.repo.Search(ctx, dto.Query, pagination)
	if err != nil {
		return nil, err
	}

	// 3. Get total count
	totalCount, err := uc.repo.SearchCount(ctx, dto.Query)
	if err != nil {
		return nil, err
	}

	// 4. Create paginated result
	result := shared.NewPaginatedResult(books, totalCount, pagination)

	// 5. Map to response
	response := application.ToPaginatedResponse(result, application.ToBookResponseList)
	return &response, nil
}
