package grpc

import (
	"context"
	"github.com/sillyhatxu/consul-client"
	"github.com/sillyhatxu/user-backend/enums"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/sillyhatxu/user-backend/model"
	"github.com/sillyhatxu/user-backend/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hv1 "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func InitialGRPC(listener net.Listener) {
	logrus.Info("---------- initial grpc server ----------")
	server := grpc.NewServer()

	healthServer := health.NewServer()
	healthServer.SetServingStatus(consul.DefaultHealthCheckGRPCServerName, hv1.HealthCheckResponse_SERVING)
	hv1.RegisterHealthServer(server, healthServer)

	user.RegisterUserServiceServer(server, &User{})
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type User struct{}

func (u *User) Add(ctx context.Context, in *user.AddRequest) (*user.Response, error) {
	err := service.Add(&model.User{
		LoginName: in.LoginName,
		Password:  in.Password,
		Channel:   in.Channel,
		Type:      in.Type,
		Status:    enums.UserStatusEnable,
	})
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Code:    user.ResponseCode_SUCCESS,
		Message: "Success",
	}, nil
}

func (u *User) Update(ctx context.Context, in *user.UpdateRequest) (*user.Response, error) {
	err := service.Update(&model.User{
		Id:        in.Id,
		LoginName: in.LoginName,
		Password:  in.Password,
		Channel:   in.Channel,
		Type:      in.Type,
		Status:    enums.UserStatusEnable,
	})
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Code:    user.ResponseCode_SUCCESS,
		Message: "Success",
	}, nil
}

func (u *User) Enable(ctx context.Context, in *user.EnableRequest) (*user.Response, error) {
	err := service.Enable(in.Id)
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Code:    user.ResponseCode_SUCCESS,
		Message: "Success",
	}, nil
}

func (u *User) Disable(ctx context.Context, in *user.DisableRequest) (*user.Response, error) {
	err := service.Disable(in.Id)
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Code:    user.ResponseCode_SUCCESS,
		Message: "Success",
	}, nil
}

func (u *User) Login(ctx context.Context, in *user.LoginRequest) (*user.Response, error) {
	err := service.Login(in.LoginName, in.Password, in.Channel, in.Type)
	if err != nil {
		//if err == responsecode.InvalidLoginNameOrPasswordError {
		//	return &user.Response{
		//		Code:    responsecode.Success,
		//		Message: responsecode.Success,
		//	}, nil
		//}
		return nil, err
	}
	return &user.Response{
		Code:    user.ResponseCode_SUCCESS,
		Message: "Success",
	}, nil
}

func (u *User) List(ctx context.Context, in *user.ListRequest) (*user.UserResponse, error) {
	panic("implement me")
}
