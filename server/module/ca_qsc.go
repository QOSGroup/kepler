package module

import "time"

type CaQsc struct {
	Id         int64  `json:"id" form:"id"`
	QosChainId string `xorm:"Varchar(50)"`
	Name       string `xorm:"Varchar(50)"`
	Csr        string `xorm:"Varchar(500)"`
	Crt        string `xorm:"Varchar(1000)"`
	ApplyId    int64
	CreateTime time.Time
	ExpireTime time.Time
}
