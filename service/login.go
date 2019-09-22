package service

import (
	"github.com/sillyhatxu/user-backend/dao"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/sillyhatxu/user-backend/responseerror"
	"github.com/sillyhatxu/user-backend/utils"
)

func Login(loginName, password string, channel user.Channel, userType user.Type) error {
	u, err := dao.FindByLoginNameAndChannelAndType(loginName, channel, userType)
	if err != nil {
		return err
	}
	if u.Password != utils.MD5(password) {
		return responseerror.InvalidLoginNameOrPasswordError
	}
	return nil
}
