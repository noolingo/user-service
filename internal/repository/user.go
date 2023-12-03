package repository

import (
	"context"
	"errors"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
)

var ErrNotFound = errors.New("user not found")

type user struct {
}

var testuser = &domain.User{
	ID:       "1",
	Name:     "test",
	Email:    "test@test.com",
	Password: "123456",
}

func (u *user) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	if id != testuser.ID {
		return nil, ErrNotFound
	}
	return testuser, nil
}

func (u *user) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if email != testuser.Email {
		return nil, ErrNotFound
	}
	return testuser, nil
}
