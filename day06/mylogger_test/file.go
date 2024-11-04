package mylogger_test

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

type Filelogger struct {
	Level       LogLevel
	filePath    string
	fileObj     *os.File
	errFileObj  *os.File
	fileName    string
	maxFileSize int64 //切割文件的大小
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *Filelogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &Filelogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //  按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *Filelogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open error,error :%v\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open error err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	return nil
}

func (f *Filelogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f *Filelogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(DEBUG) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s][%s][%s:%s:%d]%s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志大于等于ERROR级别，我还要在err日志文件中在记录一遍。
			fmt.Fprintf(f.fileObj, "[%s][%s][%s:%s:%d]%s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *Filelogger) Debug(msg string, a ...interface{}) {
	// if c.enable(DEBUG) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2) //这里传2的原因，首先这个函数是从main函数里面调用的，然后它又是通过log这个方法调用的getinfo，因此传入2.
	// fmt.Printf("[%s][DEBUG][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	f.log(DEBUG, msg, a...)
	// }
}

func (f *Filelogger) Info(msg string, a ...interface{}) {
	// if c.enable(INFO) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][INFO][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	f.log(INFO, msg, a...)
	// }
	// fmt.Println(msg)
}

func (f *Filelogger) Warning(msg string, a ...interface{}) {
	// if c.enable(WARNING) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][WARNING][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	f.log(WARNING, msg, a...)
	// }
	// fmt.Println(msg)
}
func (f *Filelogger) Error(msg string, a ...interface{}) {
	// if c.enable(ERROR) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][ERROR][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, msg)
	f.log(ERROR, msg, a...)
	// }
	// fmt.Println(msg)
}
func (f *Filelogger) Fatal(msg string, a ...interface{}) {
	// if c.enable(FATAL) {
	// now := time.Now()
	// funcName, fileName, lineNo := getInfo(2)
	// fmt.Printf("[%s][FATAL][%s:%s:%d]%s", now.Format("2006-01-02 15:04:05"), fileName, funcName, lineNumber, msg)
	f.log(FATAL, msg, a...)
	// }
	// fmt.Println(msg)
}
