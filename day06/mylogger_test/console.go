package mylogger_test

import (
	"fmt"
	"time"
)

// logger 日志结构体
type Logger struct {
	Level LogLevel
}

//Newlog 构造函数

func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func log(lv logLevel,msg string){
	fmt.Printf()
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s][DEBUG][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), ,fileName,funcName,lineNumber,msg)
	}
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
