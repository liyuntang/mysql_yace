package databaseOperator

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

func insertOperat(engine *xorm.Engine, logger *log.Logger,thread_id int)  {
	user := new(user)
	user.Name = makeName()
	user.Age = makeAge()
	user.Created = makeTime()
	_, err := engine.Insert(user)
	if err != nil {
		logger.Println("thread", thread_id, "operat is insert is bad")
	}
	logger.Println("thread", thread_id, "operat is insert, is ok")
}

func updateOperat(engine *xorm.Engine, logger *log.Logger,thread_id int)  {
	age := makeAge()
	name := makeName()
	sql := fmt.Sprintf("update user set age = %d where name = '%s'", age, name)
	fmt.Println("sql is", sql)
	_,err := engine.Exec(sql)
	if err != nil {
		logger.Println("thread", thread_id, "operat is update is bad")
	}
	logger.Println("thread", thread_id, "operat is update is ok")
}

func selectOperat(engine *xorm.Engine, logger *log.Logger,thread_id int)  {
	age := makeAge()
	name := makeName()
	sql := fmt.Sprintf("select * from user where age = '%d' or name = '%s'", age, name)
	fmt.Println("sql is", sql)
	_, err := engine.Query(sql)
	if err != nil {
		logger.Println("thread", thread_id, "operat is select is bad")
	}
	logger.Println("thread", thread_id, "operat is select is ok")
}
