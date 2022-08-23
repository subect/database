package main

import (
	"fmt"
	"time"
)

func main() {
	//ADD
	now := time.Now()
	time.Sleep(5 * time.Second)
	duration := time.Now().Sub(now)
	fmt.Println("用时：", duration)
	//later1 := now.Add(time.Hour * 1)
	//later2 := now.Add(time.Second * 5)
	//later3 := now.Add(time.Minute * 2)
	//fmt.Printf("later1:%v,later2:%v,later3:%v\n", later1, later2, later3)
}
