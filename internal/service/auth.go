package service

import (
	"context"
	"errors"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/MelnikovNA/noolingo-user-service/internal/pkg/tokens"
	"github.com/MelnikovNA/noolingo-user-service/internal/repository"
	"github.com/MelnikovNA/noolingoproto/codegen/go/apierrors"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	logger       *logrus.Logger
	config       *domain.Config
	repository   repository.Repository
	accessToken  tokens.JWTToken
	refreshToken tokens.JWTToken
}

func (a *AuthService) makeToken(ctx context.Context, user *domain.User) (accessToken string, refreshToken string, err error) {
	accessToken, err = a.accessToken.NewToken(user.ID, a.config.Auth.AccessKeyTtl)
	if err != nil {
		a.logger.WithError(err).Errorf("error generating access token")
		return "", "", apierrors.ErrInternalServerError
	}
	refreshToken, err = a.refreshToken.NewToken(user.ID, a.config.Auth.RefreshKeyTtl)
	if err != nil {
		a.logger.WithError(err).Errorf("error generating refresh token")
	}
	return accessToken, refreshToken, err
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	ParsedUserID, err := a.refreshToken.ParseToken(refreshToken)
	if err != nil {
		a.logger.WithError(err).Warn("token parse error")
		return "", "", apierrors.ErrTokenExpired
	}
	user, err := a.repository.GetUserByID(ctx, ParsedUserID)
	if err != nil {
		a.logger.WithError(err).Errorf("error in db")
		return "", "", errors.New("error in DB")
	}
	newAccessToken, newRefreshToken, err = a.makeToken(ctx, user)

	return
}
