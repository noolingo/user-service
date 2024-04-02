package tokens

import (
	"errors"
	"time"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	config *domain.Config
}

func (t *JWTToken) NewToken(userID string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"ttl":    time.Now().Add(ttl),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(t.config.Auth.AccessSecretKey)
}

func (t *JWTToken) ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return t.config.Auth.AccessSecretKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("invalid user id")
	}
	return userID, err
}
