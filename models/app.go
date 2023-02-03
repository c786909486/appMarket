package models

import "appMarket/datasource"

type AppModel struct {
	AppId       int    `json:"appId" xorm:"not null pk autoincr comment INT(11)"`
	AppName     string `json:"appName" xorm:"not null comment('app名称') VARCHAR(50)"`
	PackageName string `json:"packName" xorm:"not null comment('app包名') VARCHAR(50)"`
	AppInfo     string `json:"appInfo" xorm:"comme nt('app介绍') VARCHAR(50)"`
	AppVersion  string `json:"appVersion" xorm:"not null comment('app版本号') VARCHAR(50)"`
	GroupId     int    `json:"groupId" xorm:"not null comment('渠道id') INT(11)"`
	AppUrl      string `json:"appUrl" xorm:"not null comment('app地址') VARCHAR(50)"`
	FileMd5     string `json:"-" xorm:"not null comment('文件md5') VARCHAR(50)"`
	CreateTime  int64  `json:"createTime" xorm:"created"`
	UpdateTime  int64  `json:"updateTime" xorm:"updated"`
}

type AppUpdateModel struct {
	RecordId      int    `json:"recordId" xorm:"not null pk autoincr comment INT(11)"`
	CreateTime    int64  `json:"createTime" xorm:"created"`
	UpdateInfo    string `json:"updateInfo" xorm:"not null comment('更新内容') VARCHAR(255)"`
	UpdateVersion string `json:"updateVersion" xorm:"not null comment('更新版本号')"`
	AppId         int    `json:"appId" xorm:"not null comment('appId')"`
	UpdateTime    int64  `json:"-" xorm:"updated"`
}

func InitApp() {
	x := datasource.InstanceMaster()
	x.Sync(new(AppModel))
}
