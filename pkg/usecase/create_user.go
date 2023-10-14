package usecase

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
)

type UserCreator struct {
	repo repository.UserRepository

	reader *UserReader
}

func NewUserCreator(r repository.UserRepository, reader *UserReader) *UserCreator {
	return &UserCreator{repo: r, reader: reader}
}

func (uc *UserCreator) Create(ctx context.Context, username string, password []byte) error {
	// Check if user already exists.
	if ex, err := uc.reader.ReadByName(ctx, username); ex != nil {
		return nil // TODO: return error
	} else if err != nil {
		return err
	}

	usr := domain.NewUser(username, password)

	// Create user.
	if err := uc.repo.Create(ctx, &usr); err != nil {
		return err
	}

	return nil
}
