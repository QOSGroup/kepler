package module

import "time"

type ApplyQsc struct {
	Id         int64
	QscName    string    `xorm:"Varchar(50)" json:"qsc_name" form:"qsc_name"`
	QosChainId string    `xorm:"Varchar(50)" json:"qos_chain_id" form:"qos_chain_id"`
	QscPub     string    `xorm:"Varchar(100)" json:"qsc_pub" form:"qsc_pub"`
	BankerPub  string    `xorm:"Varchar(100)" json:"banker_pub" form:"banker_pub"`
	Email      string    `xorm:"Varchar(100)" json:"email" form:"email"`
	Phone      string    `xorm:"Varchar(20)" json:"phone" form:"phone"`
	Info       string    `xorm:"Varchar(500)" json:"info" form:"info"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Status     int       `json:"status" form:"status"`
	Note       string    `xorm:"Varchar(100)" json:"note" form:"note"`
}
