package main

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[pprof]", log.Lshortfile|log.Lmicroseconds)
	logFile, err := os.OpenFile("../xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	logger.SetOutput(logFile)
}

func main() {

	logger.Println("这是自定义的logger记录的日志。")
}
