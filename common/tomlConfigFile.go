package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"sync"
)

var (
	conf *YACE
	once sync.Once
)

func TomlConfig(configFile *string) *YACE {
	// 获取配置文件的绝对路径
	fileAbsPath := getAbsPath(configFile)

	// 采用单例模式解析配置文件
	once.Do(func() {
		_, err := toml.DecodeFile(fileAbsPath, &conf)
		if err != nil {
			fmt.Println("sorry, toml configFile", *configFile, "is bad")
			os.Exit(1)
		}
	})
	return conf
}
