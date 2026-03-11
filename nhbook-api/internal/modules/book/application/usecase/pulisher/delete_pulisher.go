package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type DeletePublisherUsecase struct {
	repo domain.IPublisherRepository
}

func NewDeletePublisherUsecase(repo domain.IPublisherRepository) *DeletePublisherUsecase {
	return &DeletePublisherUsecase{repo: repo}
}

func (uc *DeletePublisherUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	// 1. Check if publisher exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrPublisherNotFound
	}

	// 2. Delete from repository
	return uc.repo.Delete(ctx, id)
}
