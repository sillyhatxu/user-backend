package dao

import (
	"github.com/sillyhatxu/mysql-client"
	"github.com/sillyhatxu/user-backend/enums"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/sillyhatxu/user-backend/model"
	"github.com/sillyhatxu/user-backend/utils"
	"github.com/sirupsen/logrus"
)

var client *dbclient.MysqlClient

func Initial(dataSourceName, ddlPath string) error {
	mysqlClient, err := dbclient.NewMysqlClient(dataSourceName, dbclient.DDLPath(ddlPath))
	client = mysqlClient
	return err
}

const insertSQL = `
INSERT INTO user (login_name, password, channel, type, status) VALUES (?, ?, ?, ?, ?)
`

func Insert(u *model.User) error {
	logrus.Infof("insert user : %#v", u)
	_, err := client.Insert(insertSQL, u.LoginName, utils.MD5(u.Password), u.Channel, u.Type, u.Status)
	return err
}

const updateSQL = `
	UPDATE user
	SET login_name         = ?,
	    password           = ?,
	    channel            = ?,
	    type               = ?,
	    status             = ?,
	    last_modified_time = now(3)
	WHERE id = ?
`

func Update(u *model.User) error {
	logrus.Infof("update user : %#v", u)
	_, err := client.Update(updateSQL, u.LoginName, utils.MD5(u.Password), u.Channel, u.Type, u.Status, u.Id)
	return err
}

const disableSQL = `
	UPDATE user
	SET status             = ?,
	    last_modified_time = now(3)
	WHERE id = ?
`

func Disable(id int64) error {
	logrus.Infof("disable user : %d", id)
	_, err := client.Update(disableSQL, enums.UserStatusDisable, id)
	return err
}

const enableSQL = `
	UPDATE user
	SET status             = ?,
	    last_modified_time = now(3)
	WHERE id = ?
`

func Enable(id int64) error {
	logrus.Infof("enable user : %d", id)
	_, err := client.Update(enableSQL, enums.UserStatusEnable, id)
	return err
}

const findByLoginNameAndChannelAndTypeSQL = `
	SELECT *
	FROM user
	WHERE login_name = ?, password = ?, channel = ?, type = ?
`

func FindByLoginNameAndChannelAndType(loginName string, channel user.Channel, userType user.Type) (*model.User, error) {
	var u model.User
	err := client.FindFirst(findByLoginNameAndChannelAndTypeSQL, &u, loginName, channel, userType)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

const findListSQL = `
	SELECT * FROM user
`

func FindList() ([]user.User, error) {
	var array []user.User
	err := client.FindList(findListSQL, &array)
	if err != nil {
		return nil, err
	}
	return array, nil
}
