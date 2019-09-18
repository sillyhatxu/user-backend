package service

import "github.com/sillyhatxu/user-backend/dao"

func Enable(id int64) error {
	return dao.Enable(id)
}
