package service

import (
	"context"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/MelnikovNA/noolingo-user-service/internal/repository"
)

type userService struct {
	repository repository.Repository
}

func (u *userService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return u.repository.GetUserByID(ctx, id)
}
