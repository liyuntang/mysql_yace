package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"sync"
)

type database struct {
	Address string	`toml:"address"`
	Port int	`toml:"port"`
	User string	`toml:"user"`
	Passwd string	`toml:"passwd"`
	Schema string	`toml:"schema"`
	Charset string	`toml:"charset"`
}

type system struct {
	LogFile string	`toml:"logfile"`
	Rows int 	`toml:"rows"`
	Thread int `toml"thread"`
	Times int `toml:"times"`
}


type TomlConfigStruct struct {
	Database database
	System system
}

var (
	conf *TomlConfigStruct
	once sync.Once
)

func TomlConfig(configFile string) *TomlConfigStruct {
	// 判断是否是绝对路径，如果不是则获取绝对路径
	File := getAbsPath(configFile)
	// 采用单例模式解析配置
	once.Do(func() {
		_, err := toml.DecodeFile(File, &conf)
		if err != nil {
			fmt.Println("toml configFile", File, "is bad")
			os.Exit(1)
		}

	})
	return conf
}