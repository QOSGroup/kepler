package module

import "time"

type ApplyQcp struct {
	Id         int64
	QcpChainId string    `xorm:"Varchar(50)" json:"qcp_chain_id" form:"qcp_chain_id"`
	QosChainId string    `xorm:"Varchar(50)" json:"qos_chain_id" form:"qos_chain_id"`
	QcpPub     string    `xorm:"Varchar(100)" json:"qcp_pub" form:"qcp_pub"`
	Email      string    `xorm:"Varchar(100)" json:"email" form:"email"`
	Phone      string    `xorm:"Varchar(20)" json:"phone" form:"phone"`
	Info       string    `xorm:"Varchar(500)" json:"info" form:"info"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     int       `json:"status" form:"status"`
	Note       string    `xorm:"Varchar(100)" json:"note" form:"note"`
}
