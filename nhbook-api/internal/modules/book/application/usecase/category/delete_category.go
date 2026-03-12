package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type DeleteCategoryUsecase struct {
	repo domain.ICategoryRepository
}

func NewDeleteCategoryUsecase(repo domain.ICategoryRepository) *DeleteCategoryUsecase {
	return &DeleteCategoryUsecase{repo: repo}
}

func (uc *DeleteCategoryUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	// 1. Check if category exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrCategoryNotFound
	}

	// 2. Delete from repository
	return uc.repo.Delete(ctx, id)
}
