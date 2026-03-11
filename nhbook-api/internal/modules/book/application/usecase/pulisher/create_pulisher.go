package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
)

type CreatePublisherUsecase struct {
	repo domain.IPublisherRepository
}

func NewCreatePublisherUsecase(repo domain.IPublisherRepository) *CreatePublisherUsecase {
	return &CreatePublisherUsecase{repo: repo}
}

func (uc *CreatePublisherUsecase) Execute(ctx context.Context, dto *application.CreatePublisherDTO) (*application.PublisherResponseDTO, error) {
	// 1. Create domain entity
	publisher, err := domain.NewPublisher(dto.Name)
	if err != nil {
		return nil, err
	}

	// 2. Save to repository
	if err := uc.repo.Create(ctx, publisher); err != nil {
		return nil, err
	}

	// 3. Return response
	return application.ToPublisherResponse(publisher), nil
}
