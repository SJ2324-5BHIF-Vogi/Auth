package usecase

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
	"github.com/google/uuid"
)

type UserReader struct {
	repo repository.UserRepository
}

func NewUserReader(r repository.UserRepository) *UserReader {
	return &UserReader{repo: r}
}

func (ur *UserReader) Read(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	usr, err := ur.repo.Read(ctx, id)

	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (ur *UserReader) ReadByName(ctx context.Context, username string) (*domain.User, error) {
	usr, err := ur.repo.ReadByName(ctx, username)

	if err != nil {
		return nil, err
	}

	return usr, nil
}
