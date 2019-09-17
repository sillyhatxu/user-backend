package model

import (
	"github.com/sillyhatxu/user-backend/grpc/user"
	"time"
)

type User struct {
	Id               int64        `json:"id"`
	LoginName        string       `json:"login_name"`
	Password         string       `json:"password"`
	Channel          user.Channel `json:"channel"`
	Type             user.Type    `json:"type"`
	Status           bool         `json:"status"`
	LastModifiedTime time.Time    `json:"last_modified_time"`
	CreatedTime      time.Time    `json:"created_time"`
}
