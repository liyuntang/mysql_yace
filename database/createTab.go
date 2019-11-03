package database

import (
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

func createTab(engine *xorm.Engine, logger *log.Logger) {
	err := engine.Sync2(new(user))
	if err != nil {
		logger.Println("create table is bad")
		os.Exit(1)
	}
	logger.Println("create table is ok")
}
