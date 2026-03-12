package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type UpdatePublisherUsecase struct {
	repo domain.IPublisherRepository
}

func NewUpdatePublisherUsecase(repo domain.IPublisherRepository) *UpdatePublisherUsecase {
	return &UpdatePublisherUsecase{repo: repo}
}

func (uc *UpdatePublisherUsecase) Execute(ctx context.Context, id uuid.UUID, dto *application.UpdatePublisherDTO) (*application.PublisherResponseDTO, error) {
	// 1. Get existing publisher
	publisher, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if publisher == nil {
		return nil, domain.ErrPublisherNotFound
	}

	// 2. Update domain entity
	if err := publisher.Update(dto.Name); err != nil {
		return nil, err
	}

	// 3. Save to repository
	if err := uc.repo.Update(ctx, publisher); err != nil {
		return nil, err
	}

	// 4. Return response
	return application.ToPublisherResponse(publisher), nil
}
