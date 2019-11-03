package common

type database struct {
	Address string	`toml:"address"`
	Port int	`toml:"port"`
	User string	`toml:"user"`
	Passwd string	`toml:"passwd"`
	Schema string 	`toml:"schema"`
	Charset string 	`toml:"charset"`
}


type system struct {
	Logfile string	`toml:"logfile"`
	Rows int	`toml:"rows"`
	Thread int	`toml:"thread"`
	Times int	`toml:"times"`
}


type YACE struct {
	Database database
	System system
}



















































