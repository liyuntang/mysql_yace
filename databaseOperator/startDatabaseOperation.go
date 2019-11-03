package databaseOperator

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/go-xorm/xorm"
	"log"
	"math/rand"
	"os"
	"tidb_yace/common"
	"time"
)

func StartDabaseOperation(tomlConfig *common.TomlConfigStruct, logger *log.Logger) {
	// 初始化dbConfig
	dbParamet := initDbConfig(tomlConfig)
	// 初始化压测参数
	yaceParam := initYace(tomlConfig)
	// 获取engine
	engine := getEngine(dbParamet, logger)
	// 创建数据表
	//createTab(engine, logger)
	// 根据压测参数在测试表中造数据
	fmt.Println("根据压测参数在测试表中造数据")
	for thread_id := 1; thread_id <= yaceParam.thread; thread_id++ {
		engine_thread_id := getEngine(dbParamet, logger)
		go makeData(engine_thread_id, logger, yaceParam)
	}
	// 从这个地方开始进入压测阶段，压测操作在insert、update、select中随机选择
	for times := 1; times <= yaceParam.times; times++ {
		for thread_id := 1; thread_id <= yaceParam.thread; thread_id++ {
			var operation = []func(engine *xorm.Engine, logger *log.Logger,thread_id int){insertOperat, updateOperat, selectOperat}
			index := rand.Intn(3)
			startTime := time.Now()
			operation[index](engine, logger, thread_id)
			runTime := time.Since(startTime).Nanoseconds()
			logger.Println("thread is",thread_id, "runTime is",runTime)
		}


	}

}
























// 创建数据表
func createTab(engine *xorm.Engine, logger *log.Logger){
	err := engine.Sync2(new(user))
	if err != nil {
		logger.Println("create table user is bad")
		os.Exit(1)
	}
	log.Println("create table user is ok")
}



// 初始化dbConfig
func initDbConfig(tomlConfig *common.TomlConfigStruct) (DB *dbConfig) {
	//获取数据库连接参数
	db := new(dbConfig)
	db.address = tomlConfig.Database.Address
	db.port = tomlConfig.Database.Port
	db.user = tomlConfig.Database.User
	db.passwd = tomlConfig.Database.Passwd
	db.schema = tomlConfig.Database.Schema
	db.charset = tomlConfig.Database.Charset
	return db
}

func initYace(tomlConfig *common.TomlConfigStruct) (yaceParam *yace) {
	yaceP := new(yace)
	yaceP.thread = tomlConfig.System.Thread
	yaceP.times = tomlConfig.System.Times
	yaceP.rows = tomlConfig.System.Rows
	return yaceP
}

func getEngine(dbParamet *dbConfig, logger *log.Logger) (dbEngine *xorm.Engine) {
	dbEndpoint := fmt.Sprintf("%s:%d", dbParamet.address, dbParamet.port)
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", dbParamet.user, dbParamet.passwd, dbEndpoint, dbParamet.schema, dbParamet.charset)
	engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		logger.Println("connect db is bad",err)
		os.Exit(1)
	}
	fmt.Println("connect db is ok")
	return  engine
}














