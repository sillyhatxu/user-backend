package grpcerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var InvalidLoginNameOrPasswordError = status.Error(codes.InvalidArgument, "invalid user name or password")
var DuplicateUserError = status.Error(codes.AlreadyExists, "duplicate user")
