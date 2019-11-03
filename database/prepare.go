package database

import (
	"log"
	"tidb_yace/common"
)

func prepare(paramet *common.YACE, logger *log.Logger) {
	// 获取数据库连接
	engine := getEngine(paramet, logger)
	// 创建数据表
	createTab(engine, logger)
	// 根据定义的表的大小（rows）灌入数据

}
