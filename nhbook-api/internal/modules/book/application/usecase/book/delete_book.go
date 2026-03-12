package usecase

import (
	"context"

	"github.com/NewHorizonIT/nhbook-api/internal/modules/book/domain"
	"github.com/google/uuid"
)

type DeleteBookUsecase struct {
	repo domain.IBookRepository
}

func NewDeleteBookUsecase(repo domain.IBookRepository) *DeleteBookUsecase {
	return &DeleteBookUsecase{repo: repo}
}

func (uc *DeleteBookUsecase) Execute(ctx context.Context, id uuid.UUID) error {
	// 1. Check if book exists
	exists, err := uc.repo.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrBookNotFound
	}

	// 2. Delete from repository
	return uc.repo.Delete(ctx, id)
}
