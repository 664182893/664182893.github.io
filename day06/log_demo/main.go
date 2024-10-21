package main

import (
	"project/day06/mylogger_test"
	"time"
)

// type logger struct {
// }

// func Newlog()Logger {
// 	return Logger{}
// }

func main() {
	// fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Printf("error:%v", err)
	// 	return
	// }
	//debug
	//trace
	//info
	//warning
	//error
	//fatal

	//日志要支持开关
	//日志要有时间、行号、文件名、日志级别、日志信息
	id := 10010
	name := "lixiang"
	//日志文件要切割
	log := mylogger_test.NewLog("DEBUG")
	for {
		log.Debug("这是一条Debug log,id:%d,name:%s", id, name)
		log.Info("这是一条info log")
		log.Warning("这是一条warning log")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 3)
	}

}
