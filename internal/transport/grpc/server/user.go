package grpcserver

import (
	"context"

	"github.com/MelnikovNA/noolingo-user-service/internal/domain"
	"github.com/MelnikovNA/noolingo-user-service/internal/service"
	"github.com/MelnikovNA/noolingoproto/codegen/go/common"
	"github.com/MelnikovNA/noolingoproto/codegen/go/noolingo"
	"github.com/sirupsen/logrus"
)

type UserServer struct {
	noolingo.UnimplementedUserServer
	logger  *logrus.Logger
	service service.Service
}

func newUserServer(logger *logrus.Logger, service service.Service) UserServer {
	return UserServer{logger: logger, service: service}
}

func newResponse(err error) (*common.Response, error) {
	response := &common.Response{
		Result: err == nil,
	}

	if err != nil {
		response.Error = &common.Error{
			Error: err.Error(),
		}
	}

	return response, err
}

func (u UserServer) SignUp(_ context.Context, _ *noolingo.SignUpRequest) (*common.Response, error) {
	panic("not implemented") // TODO: Implement
}

func (u UserServer) SignIn(_ context.Context, _ *noolingo.SignInRequest) (*noolingo.SignInReply, error) {
	panic("not implemented") // TODO: Implement
}

func (u UserServer) Logout(_ context.Context, _ *noolingo.LogoutRequest) (*common.Response, error) {
	panic("not implemented") // TODO: Implement
}

func (u UserServer) GetUser(ctx context.Context, req *noolingo.GetUserRequest) (*noolingo.GetUserResponse, error) {
	user, err := u.service.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &noolingo.GetUserResponse{Result: &noolingo.UserObject{Id: user.ID, Name: user.Name, Email: user.Email}}, nil
}

func (u UserServer) UpdateUser(ctx context.Context, req *noolingo.UpdateUserRequest) (*common.Response, error) {
	err := u.service.UpdateUser(ctx, &domain.User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return newResponse(err)
}

func (u UserServer) DeleteUser(_ context.Context, _ *noolingo.DeleteUserRequest) (*common.Response, error) {
	panic("not implemented") // TODO: Implement
}

func (u UserServer) CreateUser(ctx context.Context, req *noolingo.CreateUserRequest) (*common.Response, error) {
	_, err := u.service.CreateUser(ctx, &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return newResponse(err)
}

func (u UserServer) UpdatePassword(_ context.Context, _ *noolingo.UpdatePasswordRequest) (*common.Response, error) {
	panic("not implemented") // TODO: Implement
}

func (u UserServer) Refresh(_ context.Context, _ *noolingo.RefreshRequest) (*noolingo.RefreshReply, error) {
	panic("not implemented") // TODO: Implement
}

func (u UserServer) mustEmbedUnimplementedUserServer() {
	panic("not implemented") // TODO: Implement
}
