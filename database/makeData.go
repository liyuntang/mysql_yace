package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
	"tidb_yace/common"
	"time"
)


func makeData(paramet *common.YACE, logger *log.Logger)  {
	// 计算出每个线程需要灌入多少数据
	rowsEveryThread := paramet.System.Rows / 200
	// 启动20个线程往数据表灌数据
	for thread_id:=1;thread_id<=200;thread_id++ {
		engine := getEngine(paramet, logger)
		waitgroup.Add(1)
		go haha(engine, logger, rowsEveryThread, thread_id, &waitgroup)

	}
	waitgroup.Wait()
	fmt.Println("insert table is over")
}

func haha(engine *xorm.Engine, logger *log.Logger, rowsEveryThread int, thread_id int, wait *sync.WaitGroup)  {
	for row:=1; row<= rowsEveryThread;row++ {
		startTime := time.Now()
		isok := insertTab(engine)
		runTime := time.Since(startTime).Nanoseconds() / 1e6
		if isok {
			logger.Println("thread", thread_id, "insert table is ok, ", rowsEveryThread, "rows need to insert, and now has inset", row, "rows,run time is", runTime, "ms")
		} else {
			logger.Println("thread", thread_id, "insert table is bad")
		}
	}
	wait.Done()
}



