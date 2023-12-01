package usecase

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
	"github.com/google/uuid"
)

type UserPasswordReader struct {
	repo repository.UserRepository
}

func NewUserPasswordReader(r repository.UserRepository) *UserPasswordReader {
	return &UserPasswordReader{repo: r}
}

func (ur *UserPasswordReader) Read(ctx context.Context, id uuid.UUID) ([]byte, error) {
	usr, err := ur.repo.Read(ctx, id)

	if err != nil {
		return nil, err
	}

	return usr.Password, nil
}

func (ur *UserPasswordReader) ReadByName(ctx context.Context, username string) ([]byte, error) {
	usr, err := ur.repo.ReadByName(ctx, username)

	if err != nil {
		return nil, err
	}

	return usr.Password, nil
}
