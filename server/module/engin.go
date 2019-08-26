package module

import (
	"github.com/QOSGroup/kepler/server/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var KEngine *xorm.Engine

func init() {
	dbConfig := config.DefaultDbConfig()
	var err error
	KEngine, err = xorm.NewEngine(dbConfig.Driver, dbConfig.DateSource())
	KEngine.ShowSQL(true)
	if err != nil {
		panic(err)
	}
}
