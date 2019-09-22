package grpcclient

import (
	"context"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"google.golang.org/grpc"
	"time"
)

const defaultTimeout = 10 * time.Second

type Client struct {
	address string
	config  *Config
}

func NewGRPCUserClient(address string, opts ...Option) *Client {
	//default
	config := &Config{
		timeout: defaultTimeout,
	}
	for _, opt := range opts {
		opt(config)
	}
	return &Client{
		address: address,
		config:  config,
	}
}

func (c Client) AddUser(req *user.AddRequest) (*user.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return userClient.Add(ctx, req)
}

func (c Client) UpdateUser(req *user.UpdateRequest) (*user.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return userClient.Update(ctx, req)
}

func (c Client) Enable(req *user.EnableRequest) (*user.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return userClient.Enable(ctx, req)
}

func (c Client) Disable(req *user.DisableRequest) (*user.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return userClient.Disable(ctx, req)
}

func (c Client) Login(req *user.LoginRequest) (*user.Response, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return userClient.Login(ctx, req)
}

func (c Client) List(req *user.ListRequest) (*user.UserResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.config.timeout)
	defer cancel()
	return userClient.List(ctx, req)
}
