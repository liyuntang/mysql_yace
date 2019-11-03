package common

import (
	"log"
	"os"
)

func Initlog() (logger *log.Logger) {
	// get log file
	//logFile := conf.System.Logfile
	//// open log file
	////file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	//if err != nil {
	//	fmt.Println("open or create log file", logFile, "is bad")
	//	os.Exit(1)
	//}
	//return log.New(io.MultiWriter(file, os.Stdout), "[tidb_yace] ", log.Ldate|log.Ltime|log.Lshortfile)
	return log.New(os.Stdout, "[tidb_yace] ", log.Ldate|log.Ltime|log.Lshortfile)
}
