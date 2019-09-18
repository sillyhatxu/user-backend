package service

import (
	"github.com/sillyhatxu/user-backend/dao"
)

func Disable(id int64) error {
	return dao.Disable(id)
}
