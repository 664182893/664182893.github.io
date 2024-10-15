package mylogger_test

import "runtime"

func main() {

}

func getInfo(n int) {
	pc, fine, line, ok := runtime.Caller()
}
