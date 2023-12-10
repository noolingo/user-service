package service

import (
	"context"
	"errors"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/MelnikovNA/noolingo-user-service/internal/repository"
)

type userService struct {
	repository repository.Repository
}

var (
	ErrNoUserFound = errors.New("no such user found")
)

func (u *userService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.repository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrNoUserFound
	}

	return user, err
}
