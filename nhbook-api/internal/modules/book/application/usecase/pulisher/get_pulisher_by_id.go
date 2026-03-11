package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/application"
	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type GetPublisherByIDUsecase struct {
	repo domain.IPublisherRepository
}

func NewGetPublisherByIDUsecase(repo domain.IPublisherRepository) *GetPublisherByIDUsecase {
	return &GetPublisherByIDUsecase{repo: repo}
}

func (uc *GetPublisherByIDUsecase) Execute(ctx context.Context, id uuid.UUID) (*application.PublisherResponseDTO, error) {
	// 1. Get publisher from repository
	publisher, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if publisher == nil {
		return nil, domain.ErrPublisherNotFound
	}

	// 2. Return response
	return application.ToPublisherResponse(publisher), nil
}
