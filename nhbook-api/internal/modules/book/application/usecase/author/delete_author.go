package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type DeleteAuthorUsecase struct {
	repo domain.IAuthorRepository
}

func NewDeleteAuthorUsecase(repo domain.IAuthorRepository) *DeleteAuthorUsecase {
	return &DeleteAuthorUsecase{repo: repo}
}

func (uc *DeleteAuthorUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	// 1. Check if author exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrAuthorNotFound
	}

	// 2. Delete from repository
	return uc.repo.Delete(ctx, id)
}
