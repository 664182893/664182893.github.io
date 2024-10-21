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
func (l Logger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s][%s][%s:%s:%d]%s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string, a ...interface{}) {
	if l.enable(DEBUG) {
		// now := time.Now()
		// funcName, fileName, lineNo := getInfo(2) //这里传2的原因，首先这个函数是从main函数里面调用的，然后它又是通过log这个方法调用的getinfo，因此传入2.
		// fmt.Printf("[%s][DEBUG][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
		log(DEBUG, msg, a...)
	}
}

func (l Logger) Info(msg string, a ...interface{}) {
	if l.enable(INFO) {
		// now := time.Now()
		// funcName, fileName, lineNo := getInfo(2)
		// fmt.Printf("[%s][INFO][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
		log(INFO, msg, a...)
	}
	// fmt.Println(msg)
}

func (l Logger) Warning(msg string, a ...interface{}) {
	if l.enable(WARNING) {
		// now := time.Now()
		// funcName, fileName, lineNo := getInfo(2)
		// fmt.Printf("[%s][WARNING][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
		log(WARNING, msg, a...)
	}
	// fmt.Println(msg)
}
func (l Logger) Error(msg string, a ...interface{}) {
	if l.enable(ERROR) {
		// now := time.Now()
		// funcName, fileName, lineNo := getInfo(2)
		// fmt.Printf("[%s][ERROR][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
		log(ERROR, msg, a...)
	}
	// fmt.Println(msg)
}
func (l Logger) Fatal(msg string, a ...interface{}) {
	if l.enable(FATAL) {
		// now := time.Now()
		// funcName, fileName, lineNo := getInfo(2)
		// fmt.Printf("[%s][FATAL][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNumber, msg)
		log(FATAL, msg, a...)
	}
	// fmt.Println(msg)
}
