package grpc

import (
	"context"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/sillyhatxu/user-backend/responsecode"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

const (
	address = "localhost:8080"
	timeout = 30 * time.Second
)

func TestUser_Add(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	loginName := "loginName"
	password := "password"
	channel := user.Channel_QQ
	userType := user.Type_MANAGEMENT
	response, err := userClient.Add(ctx, &user.AddRequest{
		LoginName: loginName,
		Password:  password,
		Channel:   channel,
		Type:      userType,
	})
	if err != nil {
		panic(err)
	}
	assert.NotNil(t, response)
	assert.EqualValues(t, response.Code, responsecode.Success)
}
