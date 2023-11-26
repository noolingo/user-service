package grpcserver

import (
	"github.com/MelnikovNA/noolingoproto/codegen/go/noolingo"
	"github.com/sirupsen/logrus"
)

type UserServer struct {
	noolingo.UnimplementedUserServer
	logger *logrus.Logger
}

func newUserServer(logger *logrus.Logger) UserServer {
	return UserServer{logger: logger}
}
