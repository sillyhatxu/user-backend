package dao

import (
	"github.com/sillyhatxu/mysql-client"
	"github.com/sillyhatxu/user-backend/grpc/user"
	"github.com/sillyhatxu/user-backend/model"
	"github.com/sirupsen/logrus"
)

var client *dbclient.MysqlClient

func Initial(dataSourceName string) error {
	client = dbclient.NewMysqlClientConf(dataSourceName)
	return client.Initial()
}

const insertSQL = `
INSERT INTO user (login_name, password, channel, type) VALUES (?, ?, ?, ?)
`

func Insert(u *model.User) error {
	logrus.Infof("insert user : %#v", u)
	_, err := client.Insert(insertSQL, u.LoginName, u.Password, u.Channel, u.Channel)
	return err
}

const updateSQL = `
	UPDATE user
	SET login_name         = ?,
	    password           = ?,
	    channel            = ?,
	    type               = ?,
	    last_modified_time = now(3)
	WHERE id = ?
`

func Update(u *model.User) error {
	logrus.Infof("update user : %#v", u)
	_, err := client.Update(updateSQL, u.LoginName, u.Password, u.Channel, u.Type, u.Id)
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
