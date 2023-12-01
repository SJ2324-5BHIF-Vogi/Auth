package usecase

import (
	"bytes"
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/repository"
	"github.com/google/uuid"
)

type UserUpdater struct {
	repo repository.UserRepository

	reader *UserReader

	passwordReader *UserPasswordReader
}

func NewUserUpdater(r repository.UserRepository, reader *UserReader, passwordReader *UserPasswordReader) *UserUpdater {
	return &UserUpdater{repo: r, reader: reader}
}

func (uu *UserUpdater) Update(ctx context.Context, id uuid.UUID, username string, password []byte) error {
	// Check whether username or password should be updated.
	if password == nil && username == "" {
		return nil // TODO: return error (nothing to update)
	}

	// Check if user exists.
	ex, err := uu.reader.Read(ctx, id)

	if err != nil {
		return err
	} else if ex == nil {
		return nil // TODO: return error
	}

	exp, err := uu.passwordReader.Read(ctx, id)

	// Check if password should be updated.
	if err != nil {
		return err
	} else if password == nil {
		password = exp
	} else if bytes.Equal(exp, password) { // Check if password is the same.
		return nil // TODO: return error
	}

	// Check if username should be updated.
	if username == "" {
		username = ex.Username
	}

	usr := domain.NewUser(username, password)

	// Update user.
	if err := uu.repo.Update(ctx, id, &usr); err != nil {
		return err
	}

	return nil
}
