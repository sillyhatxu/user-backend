package service

import (
	"github.com/sillyhatxu/user-backend/dao"
	"github.com/sillyhatxu/user-backend/model"
)

func Update(user *model.User) error {
	return dao.Update(user)
}
