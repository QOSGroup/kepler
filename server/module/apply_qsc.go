package module

import "time"

type ApplyQsc struct {
	Id         int64
	QscName    string    `xorm:"Varchar(50)" json:"qscName" form:"qscName"`
	QosChainId string    `xorm:"Varchar(50)" json:"qosChainId" form:"qosChainId"`
	QscPub     string    `xorm:"Varchar(100)" json:"qscPub" form:"qscPub"`
	BankerPub  string    `xorm:"Varchar(100)" json:"bankerPub" form:"bankerPub"`
	Email      string    `xorm:"Varchar(100)" json:"email" form:"email"`
	Phone      string    `xorm:"Varchar(20)" json:"phone" form:"phone"`
	Info       string    `xorm:"Varchar(500)" json:"info" form:"info"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status" form:"status"`
	Note       string    `xorm:"Varchar(100)" json:"note" form:"note"`
}
