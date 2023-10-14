package usecase

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
	"github.com/google/uuid"
)

type UserDeleter struct {
	repo repository.UserRepository

	reader *UserReader
}

func NewUserDeleter(r repository.UserRepository, reader *UserReader) *UserDeleter {
	return &UserDeleter{repo: r, reader: reader}
}

func (ud *UserDeleter) Delete(ctx context.Context, id uuid.UUID) error {
	// Check if user exists.
	if ex, err := ud.reader.Read(ctx, id); ex == nil {
		return nil // TODO: return error
	} else if err != nil {
		return err
	}

	// Delete user.
	if err := ud.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
