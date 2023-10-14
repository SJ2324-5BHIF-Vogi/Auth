package usecase

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
	"github.com/google/uuid"
)

type UserUpdater struct {
	repo repository.UserRepository

	reader *UserReader
}

func NewUserUpdater(r repository.UserRepository, reader *UserReader) *UserUpdater {
	return &UserUpdater{repo: r, reader: reader}
}

func (uu *UserUpdater) Update(ctx context.Context, id uuid.UUID, username string, password []byte) error {
	// Check if user exists.
	ex, err := uu.reader.Read(ctx, id)

	if ex == nil {
		return nil // TODO: return error
	} else if err != nil {
		return err
	}

	// Check whether username or password should be updated.
	if username == "" {
		username = ex.Username
	} else if password == nil {
		password = ex.Password
	} else if password == nil && username == "" {
		return nil // TODO: return error (nothing to update)
	}

	usr := domain.NewUser(username, password)

	// Update user.
	if err := uu.repo.Update(ctx, id, &usr); err != nil {
		return err
	}

	return nil
}
