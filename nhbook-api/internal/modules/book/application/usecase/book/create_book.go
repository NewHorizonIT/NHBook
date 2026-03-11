package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
)

type CreateBookUsecase struct {
	bookRepo      domain.IBookRepository
	authorRepo    domain.IAuthorRepository
	categoryRepo  domain.ICategoryRepository
	publisherRepo domain.IPublisherRepository
}

func NewCreateBookUsecase(
	bookRepo domain.IBookRepository,
	authorRepo domain.IAuthorRepository,
	categoryRepo domain.ICategoryRepository,
	publisherRepo domain.IPublisherRepository,
) *CreateBookUsecase {
	return &CreateBookUsecase{
		bookRepo:      bookRepo,
		authorRepo:    authorRepo,
		categoryRepo:  categoryRepo,
		publisherRepo: publisherRepo,
	}
}

func (uc *CreateBookUsecase) Execute(ctx context.Context, dto *application.CreateBookDTO) (*application.BookResponseDTO, error) {
	// 1. Validate foreign keys exist
	if dto.AuthorID != nil {
		exists, err := uc.authorRepo.Exists(ctx, *dto.AuthorID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, domain.ErrAuthorNotFound
		}
	}

	if dto.CategoryID != nil {
		exists, err := uc.categoryRepo.Exists(ctx, *dto.CategoryID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, domain.ErrCategoryNotFound
		}
	}

	if dto.PublisherID != nil {
		exists, err := uc.publisherRepo.Exists(ctx, *dto.PublisherID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, domain.ErrPublisherNotFound
		}
	}

	// 2. Create domain entity
	book, err := domain.NewBook(domain.NewBookParams{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		AuthorID:    dto.AuthorID,
		CategoryID:  dto.CategoryID,
		PublisherID: dto.PublisherID,
		CoverImage:  dto.CoverImage,
	})
	if err != nil {
		return nil, err
	}

	// 3. Save to repository
	if err := uc.bookRepo.Create(ctx, book); err != nil {
		return nil, err
	}

	// 4. Return response
	return application.ToBookResponse(book), nil
}
