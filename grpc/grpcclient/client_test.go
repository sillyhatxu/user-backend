package grpcclient

import (
	"fmt"
	"github.com/sillyhatxu/user-backend/grpc/grpcerror"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

const (
	address = "localhost:8080"
	timeout = 30 * time.Second
)

func TestUser_Add(t *testing.T) {
	client := NewGRPCUserClient(address, Timeout(timeout))
	loginName := "sillyhat"
	password := "sillyhat"
	channel := user.Channel_REGISTER
	userType := user.Type_LEARNING_ENGLISH
	_, err := client.AddUser(&user.AddRequest{
		LoginName: loginName,
		Password:  password,
		Channel:   channel,
		Type:      userType,
	})
	if err != nil {
		errorStatus := status.FromContextError(err)
		assert.EqualValues(t, errorStatus.Code(), codes.AlreadyExists)
		assert.EqualValues(t, err, grpcerror.DuplicateUserError)
	}
}

func TestUser_AddDuplicate(t *testing.T) {
	client := NewGRPCUserClient(address, Timeout(timeout))
	loginName := "sillyhat"
	password := "sillyhat"
	channel := user.Channel_REGISTER
	userType := user.Type_LEARNING_ENGLISH
	_, err := client.AddUser(&user.AddRequest{
		LoginName: loginName,
		Password:  password,
		Channel:   channel,
		Type:      userType,
	})
	if err != nil {
		errorStatus := status.FromContextError(err)
		fmt.Println(errorStatus.Code())
		fmt.Println(codes.AlreadyExists)
		assert.EqualValues(t, errorStatus.Code(), codes.AlreadyExists)
		assert.EqualValues(t, err, grpcerror.DuplicateUserError)
	}
}
