package databaseOperator

import (
	"github.com/go-xorm/xorm"
	"log"
	"math/rand"
	"os"
	"time"
)

func makeData(engine *xorm.Engine, logger *log.Logger, yaceP *yace)  {
	for row:=1;row<=yaceP.rows;row++ {
		for thread_id:=1;thread_id<=yaceP.thread;thread_id++ {
			user := new(user)
			user.Name = makeName()
			user.Age = makeAge()
			user.Created = makeTime()
			_, err := engine.Insert(user)
			if err != nil {
				logger.Println("thread", thread_id, "insert a record is bad")
				os.Exit(1)
			}
			logger.Println("thread", thread_id, "insert a record to table, all rows is", yaceP.rows, "and now has inserted",row, "records")
		}
	}
	logger.Println("makeData to table is over")
}

func makeName() (nameString string) {
	list := []string{"","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O",
		"P","Q","R","S","T","U","V","W","X","Y","Z"}
	var name string
	for i:=0;i<10;i++ {
		index :=rand.Intn(25)
		str := list[index]
		name += str
	}
	return name
}

func makeAge() int {
	return rand.Intn(100)
}

func makeTime() int64 {
	return time.Now().UnixNano()
}