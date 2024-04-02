package service

import (
	"context"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/MelnikovNA/noolingo-user-service/internal/pkg/tokens"
	"github.com/MelnikovNA/noolingoproto/codegen/go/apierrors"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	logger       *logrus.Logger
	config       *domain.Config
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
		switch err{
			case tokens.
		}
	}

	return newAccessToken, newRefreshToken, nil
}
