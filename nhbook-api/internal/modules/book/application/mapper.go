package application

import (
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/NewHorizonIT/nhbook-api/internal/shared"
)

// ============================================================================
// AUTHOR MAPPERS
// ============================================================================

func ToAuthorResponse(author *domain.Author) *AuthorResponseDTO {
	return &AuthorResponseDTO{
		ID:        author.ID,
		Name:      author.Name,
		Bio:       author.Bio,
		CreatedAt: author.CreatedAt,
	}
}

func ToAuthorResponseList(authors []domain.Author) []AuthorResponseDTO {
	result := make([]AuthorResponseDTO, len(authors))
	for i, author := range authors {
		result[i] = *ToAuthorResponse(&author)
	}
	return result
}

// ============================================================================
// CATEGORY MAPPERS
// ============================================================================

func ToCategoryResponse(category *domain.Category) *CategoryResponseDTO {
	return &CategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func ToCategoryResponseList(categories []domain.Category) []CategoryResponseDTO {
	result := make([]CategoryResponseDTO, len(categories))
	for i, category := range categories {
		result[i] = *ToCategoryResponse(&category)
	}
	return result
}

// ============================================================================
// PUBLISHER MAPPERS
// ============================================================================

func ToPublisherResponse(publisher *domain.Publisher) *PublisherResponseDTO {
	return &PublisherResponseDTO{
		ID:        publisher.ID,
		Name:      publisher.Name,
		CreatedAt: publisher.CreatedAt,
	}
}

func ToPublisherResponseList(publishers []domain.Publisher) []PublisherResponseDTO {
	result := make([]PublisherResponseDTO, len(publishers))
	for i, publisher := range publishers {
		result[i] = *ToPublisherResponse(&publisher)
	}
	return result
}

// ============================================================================
// BOOK MAPPERS
// ============================================================================

func ToBookResponse(book *domain.Book) *BookResponseDTO {
	var authorID, categoryID, publisherID *string
	if book.AuthorID != nil {
		id := book.AuthorID.String()
		authorID = &id
	}
	if book.CategoryID != nil {
		id := book.CategoryID.String()
		categoryID = &id
	}
	if book.PublisherID != nil {
		id := book.PublisherID.String()
		publisherID = &id
	}

	return &BookResponseDTO{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price.String(),
		Stock:       book.Stock.Value(),
		AuthorID:    authorID,
		CategoryID:  categoryID,
		PublisherID: publisherID,
		CoverImage:  book.CoverImage,
		IsActive:    book.IsActive,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}
}

func ToBookWithRelationsResponse(book *domain.BookWithRelations) *BookResponseDTO {
	response := ToBookResponse(&book.Book)
	response.AuthorName = book.AuthorName
	response.CategoryName = book.CategoryName
	response.PublisherName = book.PublisherName
	return response
}

func ToBookResponseList(books []domain.BookWithRelations) []BookResponseDTO {
	result := make([]BookResponseDTO, len(books))
	for i, book := range books {
		result[i] = *ToBookWithRelationsResponse(&book)
	}
	return result
}

// ============================================================================
// PAGINATION MAPPERS
// ============================================================================

func ToPaginatedResponse[T any, R any](result shared.PaginatedResult[T], mapper func([]T) []R) PaginatedResponse[R] {
	return PaginatedResponse[R]{
		Items:      mapper(result.Items),
		TotalCount: result.TotalCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	}
}
