package dao

import (
	"github.com/sillyhatxu/user-backend/enums"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/sillyhatxu/user-backend/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	err := Initial("sillyhat_user:sillyhat_user@tcp(127.0.0.1:3308)/sillyhat_user?loc=Asia%2FSingapore")
	if err != nil {
		panic(err)
	}
}

func TestInsert(t *testing.T) {
	u := &model.User{
		LoginName: "admin",
		Password:  "123",
		Channel:   user.Channel_REGISTER,
		Type:      user.Type_MANAGEMENT,
		Status:    enums.UserStatusEnable,
	}
	err := Insert(u)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	u := &model.User{
		Id:        1,
		LoginName: "sillyhat",
		Password:  "123",
		Channel:   user.Channel_INSTAGRAM,
		Type:      user.Type_WORD,
		Status:    enums.UserStatusDisable,
	}
	err := Update(u)
	assert.Nil(t, err)
}
