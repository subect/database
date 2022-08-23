package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("../xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	log.SetPrefix("[pprof]")
}

func main() {
	//log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	//log.Println("这是一条很普通的日志。")
	//log.SetPrefix("[pprof]")
	log.Println("这是一条加了前缀的日志。")
}
