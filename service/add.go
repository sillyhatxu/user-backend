package service

import (
	"github.com/sillyhatxu/user-backend/dao"
	"github.com/sillyhatxu/user-backend/model"
)

func Add(user *model.User) error {
	return dao.Insert(user)
}
