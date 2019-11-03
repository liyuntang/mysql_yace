package database

type user struct {
	Id int64
	Name string	`xorm:"varchar(25) index"`
	Age int		`xorm:"index"`
	Created int64 `xorm:"index"`
}
