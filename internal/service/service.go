package service

import (
	"context"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/MelnikovNA/noolingo-user-service/internal/repository"
)

type Service interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

func New(r repository.Repository) Service {
	return &userService{repository: r}
}
