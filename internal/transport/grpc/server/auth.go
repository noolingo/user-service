package grpcserver

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc/metadata"
)

type Requester struct {
	AccessToken  string
	UserID       string
	IsAdminValue bool
}

func (r *Requester) GetAccessToken() string {
	return r.AccessToken
}

func (r *Requester) GetUserID() string {
	return r.UserID
}

func (r *Requester) IsAdmin() bool {
	return r.IsAdminValue
}

func fromContext(ctx context.Context) (*Requester, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata provided")
	}

	accessToken, userID, isAdmin := md["access_token"], md["user_id"], md["is_admin"]

	if len(userID) == 0 || len(isAdmin) == 0 || len(accessToken) == 0 {
		return nil, errors.New("invalid metadata")
	}

	isAdminValue, err := strconv.ParseBool(isAdmin[0])
	if err != nil {
		return nil, err
	}

	return &Requester{
		AccessToken:  accessToken[0],
		UserID:       userID[0],
		IsAdminValue: isAdminValue,
	}, nil
}

func Auth(ctx context.Context) (r *Requester, err error) {
	r, err = fromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
} //
