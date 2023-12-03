package repository

import (
	"context"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
)

type Repository interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

func New() Repository {
	return &user{}
}
