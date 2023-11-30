package dto

import (
	"github.com/SJ2324-5BHIF-Vogi/Auth/pkg/domain"
	"github.com/google/uuid"
)

type UserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func ModelToDto(user *domain.User) *UserDTO {
	return &UserDTO{
		ID:       user.ID.String(),
		Username: user.Username,
	}
}

func DtoFromModel(dto *UserDTO) (*domain.User, error) {
	id, err := uuid.Parse(dto.ID)

	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       id,
		Username: dto.Username,
	}, nil
}
