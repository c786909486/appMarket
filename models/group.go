package models

import "appMarket/datasource"

//渠道model
type ChannelModel struct {
	ChannelId   string `json:"channelId" xorm:"not null pk autoincr comment INT(11)"`
	ChannelName string `json:"channelName" xorm:"not null comment('渠道名称') VARCHAR(50)"`
	CreateTime  int64  `json:"-" xorm:"created"`
	CreatorId   int    `json:"creatorId" xorm:"not null comment('创建人id') INT(11)"`
	CreatorName string `json:"creatorName" xorm:"not null comment('创建人姓名') VARCHAR(50)"`
	UpdateTime  int64  `json:"-" xorm:"updated"`
}

func InitChannel() {
	x := datasource.InstanceMaster()
	x.Sync(new(ChannelModel))
}
