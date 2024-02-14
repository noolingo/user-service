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

func (u *userService) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	return u.repository.CreateUser(ctx, user)
}

func (u *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	user2, err := u.GetUserByID(ctx, user.ID)
	if err != nil {
		return err
	}
	user.Password = user2.Password
	return u.repository.UpdateUser(ctx, user)
}

func (u *userService) DeleteUser(ctx context.Context, user *domain.User) error {
	err := u.DeleteUser(ctx, user.ID)
	if err != nil {
		return err
	}
	return err
}
