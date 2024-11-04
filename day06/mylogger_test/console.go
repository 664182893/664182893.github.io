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

// 判断日志级别
func (c Logger) enable(LogLevel LogLevel) bool {
	return c.Level <= LogLevel
}

func (c Logger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(DEBUG) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s][%s][%s:%s:%d]%s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c Logger) Debug(msg string, a ...interface{}) {
	// if c.enable(DEBUG) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2) //这里传2的原因，首先这个函数是从main函数里面调用的，然后它又是通过log这个方法调用的getinfo，因此传入2.
	// fmt.Printf("[%s][DEBUG][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	c.log(DEBUG, msg, a...)
	// }
}

func (c Logger) Info(msg string, a ...interface{}) {
	// if c.enable(INFO) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][INFO][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	c.log(INFO, msg, a...)
	// }
	// fmt.Println(msg)
}

func (c Logger) Warning(msg string, a ...interface{}) {
	// if c.enable(WARNING) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][WARNING][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	c.log(WARNING, msg, a...)
	// }
	// fmt.Println(msg)
}
func (c Logger) Error(msg string, a ...interface{}) {
	// if c.enable(ERROR) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][ERROR][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	c.log(ERROR, msg, a...)
	// }
	// fmt.Println(msg)
}
func (c Logger) Fatal(msg string, a ...interface{}) {
	// if c.enable(FATAL) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][FATAL][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNumber, msg)
	c.log(FATAL, msg, a...)
	// }
	// fmt.Println(msg)
}
