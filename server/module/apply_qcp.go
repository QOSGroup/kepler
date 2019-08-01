package module

import "time"

type ApplyQcp struct {
	Id         int64
	QcpChainId string    `xorm:"Varchar(50)" json:"qcpChainId" form:"qcpChainId"`
	QosChainId string    `xorm:"Varchar(50)" json:"qosChainId" form:"qosChainId"`
	QcpPub     string    `xorm:"Varchar(100)" json:"qcpPub" form:"qcpPub"`
	Email      string    `xorm:"Varchar(100)" json:"email" form:"email"`
	Phone      string    `xorm:"Varchar(20)" json:"phone" form:"phone"`
	Info       string    `xorm:"Varchar(500)" json:"info" form:"info"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status     int       `json:"status" form:"status"`
	Note       string    `xorm:"Varchar(100)" json:"note" form:"note"`
}
