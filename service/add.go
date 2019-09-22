package service

import (
	"github.com/sillyhatxu/user-backend/dao"
	"github.com/sillyhatxu/user-backend/grpc/grpcerror"
	"github.com/sillyhatxu/user-backend/model"
)

func Add(user *model.User) error {
	u, err := dao.FindByLoginNameAndChannelAndType(user.LoginName, user.Channel, user.Type)
	if err != nil {
		return err
	}
	if u != nil {
		return grpcerror.DuplicateUserError
	}
	return dao.Insert(user)
}
