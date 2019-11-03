package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"tidb_yace/common"
)

func getEngine(paramet *common.YACE, logger *log.Logger) (engine *xorm.Engine) {
	endPoint := fmt.Sprintf("%s:%d", paramet.Database.Address, paramet.Database.Port)
	user := paramet.Database.User
	passwd := paramet.Database.Passwd
	schema := paramet.Database.Schema
	charset := paramet.Database.Charset
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", user, passwd, endPoint, schema, charset)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		logger.Println("init db connection is bad")
		os.Exit(1)
	}
	return engine
}
