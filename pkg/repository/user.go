package repository

import (
	"context"

	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error

	Read(ctx context.Context, id uuid.UUID) (*domain.User, error)

	ReadByName(ctx context.Context, username string) (*domain.User, error)

	Update(ctx context.Context, id uuid.UUID, user *domain.User) error

	Delete(ctx context.Context, id uuid.UUID) error
}
