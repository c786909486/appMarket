package models

import (
	"appMarket/datasource"
	"appMarket/myUtils"
)

type UserModel struct {
	UserId       int    `json:"userId" xorm:"not null pk autoincr comment INT(11)"`
	UserName     string `json:"userName" xorm:"not null comment('用户名称') VARCHAR(50)"`
	UserAccount  string `json:"userAccount" xorm:"not null comment('用户账号') VARCHAR(50)"`
	UserPassword string `json:"-" xorm:"not null comment('用户密码') VARCHAR(255)"`
	UserHead     string `json:"userHead" xorm:"comment('用户头像') VARCHAR(255)"`
	CreateTime   int64  `json:"-" xorm:"created"`
	UpdateTime   int64  `json:"-" xorm:"updated"`
	Version      int    `json:"-" xorm:"version"`
}

func InitUser() {
	x := datasource.InstanceMaster()
	x.Sync(new(UserModel))
	users, _ := x.Query("select * from user_model")
	if len(users) == 0 {
		defaultUser := new(UserModel)
		defaultUser.UserAccount = "admin"
		defaultUser.UserName = "管理员"
		defaultUser.UserPassword = myUtils.Md5V2("123456")
		x.Insert(defaultUser)
	}
}
