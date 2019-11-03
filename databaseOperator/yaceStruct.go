package databaseOperator

type dbConfig struct {
	address string
	port int
	user string
	passwd string
	schema string
	charset string
}


// 压测表
type user struct {
	Id int64
	Name string	`xorm:"varchar(25) index"`
	Age int		`xorm:"index"`
	Created int64 `xorm:"index"`
}

type yace struct {
	thread int
	times int
	rows int
}