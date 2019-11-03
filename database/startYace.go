package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"math/rand"
	"sync"
	"tidb_yace/common"
	"time"
)
var (
	waitgroup sync.WaitGroup
)
func StartYace(paramet *common.YACE, logger *log.Logger, ch chan map[string]float64)  {
	// 准备工作---创建数据表
	//prepare(paramet, logger)
	// 准备工作---写入数据
	//makeData(paramet, logger)
	// 开始压测
	// 计算出每个线程需要压测多少次
	timesEveryThread := paramet.System.Times / paramet.System.Thread
	// 按照线程数启动线程进行压测
	for thread_id:=1;thread_id<=paramet.System.Thread;thread_id++ {
		engine := getEngine(paramet, logger)
		waitgroup.Add(1)
		go goOn(engine, logger, timesEveryThread, thread_id, ch, &waitgroup)

	}
	waitgroup.Wait()
	close(ch)
	fmt.Println("压测完成")
	// 清理压测数据
}

func goOn(engine *xorm.Engine, logger *log.Logger, timesEveryThread int, thread_id int, channel chan map[string]float64, wait *sync.WaitGroup)  {
	list := []string{"insertTab", "updateTabByName", "updateTabByAge", "selectTabByName", "selectTabByAge"}
	tableOperatorSlice := []func(engine *xorm.Engine) (isok bool) {insertTab, updateTabByName, updateTabByAge, selectTabByName, selectTabByAge}
	for times:=1; times<= timesEveryThread;times++ {
		startTime := time.Now()
		index := rand.Intn(len(tableOperatorSlice))
		operator := tableOperatorSlice[index]
		isok := operator(engine)
		// Nanoseconds 纳秒
		runTime := time.Since(startTime).Nanoseconds() / 1e6
		runTimeFloat64 := transfer(runTime, logger)
		if isok {
			logger.Println("thread", thread_id, list[index], "is ok, ", timesEveryThread, "times need to run, and now has run", times, "times,run time is", runTimeFloat64, "ms")
			// 向channl记录压测结果
			dict := map[string]float64{list[index]:runTimeFloat64}
			channel <- dict
		} else {
			logger.Println("thread", thread_id, list[index], "is bad")
			dict := map[string]float64{list[index]:0.00}
			channel <- dict
		}
	}
	wait.Done()
}
