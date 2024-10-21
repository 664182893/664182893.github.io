package mylogger_test

import (
	"fmt"
	"path"
	"runtime"
)

func main() {

}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FunForPC(pc).Name()
	fileName = path.Base(file)
	return
}
