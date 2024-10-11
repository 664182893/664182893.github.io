package mylogger_test

import "fmt"

//logger 日志结构体
type Logger struct {
}

//Newlog 构造函数

func NewLog() Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
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
