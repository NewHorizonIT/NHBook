package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type UpdateBookUsecase struct {
	bookRepo      domain.IBookRepository
	authorRepo    domain.IAuthorRepository
	categoryRepo  domain.ICategoryRepository
	publisherRepo domain.IPublisherRepository
}

func NewUpdateBookUsecase(
	bookRepo domain.IBookRepository,
	authorRepo domain.IAuthorRepository,
	categoryRepo domain.ICategoryRepository,
	publisherRepo domain.IPublisherRepository,
) *UpdateBookUsecase {
	return &UpdateBookUsecase{
		bookRepo:      bookRepo,
		authorRepo:    authorRepo,
		categoryRepo:  categoryRepo,
		publisherRepo: publisherRepo,
	}
}

func (uc *UpdateBookUsecase) Execute(ctx context.Context, id uuid.UUID, dto *application.UpdateBookDTO) (*application.BookResponseDTO, error) {
	// 1. Get existing book
	book, err := uc.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, domain.ErrBookNotFound
	}

	// 2. Validate foreign keys if being updated
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

	// 3. Update domain entity
	if err := book.Update(domain.UpdateBookParams{
		Title:       dto.Title,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		AuthorID:    dto.AuthorID,
		CategoryID:  dto.CategoryID,
		PublisherID: dto.PublisherID,
		CoverImage:  dto.CoverImage,
	}); err != nil {
		return nil, err
	}

	// 4. Save to repository
	if err := uc.bookRepo.Update(ctx, book); err != nil {
		return nil, err
	}

	// 5. Return response
	return application.ToBookResponse(book), nil
}
