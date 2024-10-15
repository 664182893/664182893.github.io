package mylogger_test

import "fmt"

//logger 日志结构体
type Logger struct {
	Level LogLevel
}

//Newlog 构造函数

func NewLog(levelStr string) Logger {
	level,err :=parseLogLevel(levelStr)
	if err !=nil{
		panic(err)
	}
	return Logger{
		Level:level,
	}
}

func (l Logger) Debug(msg string) {
	now :=time.Now()
	fmt.Printf("[%s] %s" now.Format("2006-01-02 15:04:05"),msg)
}
func (l Logger) Info(msg string) {
	fmt.Println(msg) 
}
func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}
func (l Logger) Error(msg string) {
	fmt.Println(msg)
}
func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}
