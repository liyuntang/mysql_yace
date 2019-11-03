package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// 判断config参数的值的类型（相对路径、绝对路径）如果是绝对路径则直接传给解析函数，如果是相对路径则转化成绝对路径
func getAbsPath(configFile string) (fileAbsPath string) {
	file, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println("get abs path is bad")
		os.Exit(1)
	}
	return file
}
