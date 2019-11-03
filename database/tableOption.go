package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"math/rand"
	"time"
)

func insertTab(engine *xorm.Engine) (isok bool) {
	name := makeName()
	age := makeAge()
	Time := makeTime()
	sql := fmt.Sprintf("insert into user(name, age, created) values ('%s', %d, %d)", name, age, Time)
	_, err := engine.Exec(sql)
	if err != nil {
		return false
	}
	return true
}

func updateTabByName(engine *xorm.Engine) (isok bool) {
	name := makeName()
	age := makeAge()
	sql := fmt.Sprintf("update user set age = %d where name = '%s' limit 1000", age, name)
	_, err := engine.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func updateTabByAge(engine *xorm.Engine) (isok bool) {
	name := makeName()
	age := makeAge()
	sql := fmt.Sprintf("update user set name = '%s' where age = %d limit 1000", name, age)
	_, err := engine.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func selectTabByName(engine *xorm.Engine) (isok bool) {
	name := makeName()
	sql := fmt.Sprintf("select sum(age), count(name), max(age), max(created) from user where name = '%s'", name)
	_, err := engine.Query(sql)
	if err != nil {
		return false
	}
	return true
}

func selectTabByAge(engine *xorm.Engine) (isok bool) {
	age:= makeAge()
	sql := fmt.Sprintf("select sum(age), count(name), max(age), max(created) from user where age = %d", age)
	_, err := engine.Query(sql)
	if err != nil {
		return false
	}
	return true
}

func makeName() (nameString string) {
	list := []string{"","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O",
		"P","Q","R","S","T","U","V","W","X","Y","Z"}
	var name string
	for i:=0;i<20;i++ {
		index :=rand.Intn(25)
		str := list[index]
		name += str
	}
	return name
}

func makeAge() (ageInt int) {
	return rand.Intn(100)
}

func makeTime() (timeInt64 int64) {
	return time.Now().UnixNano()
}
