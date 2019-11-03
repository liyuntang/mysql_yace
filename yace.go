package main

import (
	"flag"
	"fmt"
	"tidb_yace/common"
	"tidb_yace/database"
	"tidb_yace/result"
	"time"
)

var configFile = flag.String("config", "config/yace.conf", "please tell me your configration file")

func main()  {
	flag.Parse()
	// 解析配置文件
	configration := common.TomlConfig(configFile)
	// 初始化日志接口
	logger := common.Initlog()

	// 开始压测，该过程需要放到后台执行
	// 声明一个channel接收压测结果
	resultChannel := make(chan map[string]float64, 10000000)
	startTime := time.Now()
	database.StartYace(configration, logger, resultChannel)
	runTime := time.Since(startTime).Nanoseconds() / 1e6
	fmt.Println("runTime is", runTime)
	// 分析压测结果（每10秒分析、总分析）
	result.GetResult(resultChannel, runTime)
	//fmt.Println("select is", len(resultSlice.SelectOptionSlice))
	//fmt.Println("insert is", len(resultSlice.InsertOptionSlice))
	//fmt.Println("update is", len(resultSlice.UpdateOptionSlice))
	//fmt.Println("error is", len(resultSlice.ErrorOptionSlice))

}












