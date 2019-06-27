package module

import "time"

const (
	ROOT     int = 1
	ROOT_QSC int = 2
	ROOT_QCP int = 3
)

type RootCa struct {
	Id         int64
	ChainId    string `xorm:"Varchar(50)"`
	PrivKey    string `xorm:"Varchar(150)"`
	PubKey     string `xorm:"Varchar(100)"`
	Type       int
	CreateTime time.Time
}
