package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
)

type GetBooksUsecase struct {
	repo domain.IBookRepository
}

func NewGetBooksUsecase(repo domain.IBookRepository) *GetBooksUsecase {
	return &GetBooksUsecase{repo: repo}
}

func (uc *GetBooksUsecase) Execute(ctx context.Context, dto *application.GetBooksDTO) (*application.PaginatedResponse[application.BookResponseDTO], error) {
	// 1. Create pagination
	pagination := shared.NewPagination(dto.Page, dto.PageSize)

	var books []domain.BookWithRelations
	var totalCount int64
	var err error

	// 2. Get books based on filters
	if dto.CategoryID != nil {
		books, err = uc.repo.GetByCategory(ctx, *dto.CategoryID, pagination)
	} else if dto.AuthorID != nil {
		books, err = uc.repo.GetByAuthor(ctx, *dto.AuthorID, pagination)
	} else if dto.ActiveOnly {
		books, err = uc.repo.GetActive(ctx, pagination)
		if err == nil {
			totalCount, err = uc.repo.GetActiveCount(ctx)
		}
	} else {
		books, err = uc.repo.GetAll(ctx, pagination)
		if err == nil {
			totalCount, err = uc.repo.GetCount(ctx)
		}
	}

	if err != nil {
		return nil, err
	}

	// 3. Get total count if not already fetched
	if totalCount == 0 && len(books) > 0 {
		totalCount, err = uc.repo.GetCount(ctx)
		if err != nil {
			return nil, err
		}
	}

	// 4. Create paginated result
	result := shared.NewPaginatedResult(books, totalCount, pagination)

	// 5. Map to response
	response := application.ToPaginatedResponse(result, application.ToBookResponseList)
	return &response, nil
}
