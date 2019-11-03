package common

import (
	"fmt"
	"os"
	"path/filepath"
)

// 判读传递的文件路径是否是绝对路径，不过不是则获取绝地路径，并返回
func getAbsPath(file *string) (fileAbsPath string) {
	// 判断file是否是绝对路径，如果是则直接返回，如果不是则获取绝对路径，并返回
	fileAbsPath, err := filepath.Abs(*file)
	if err != nil {
		fmt.Println("get file", *file, "abs path is bad")
		os.Exit(1)
	}
	return fileAbsPath
}


