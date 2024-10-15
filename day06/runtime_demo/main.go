package main

import (
	"fmt"
	"runtime"
)

func f() {
	pc, fine, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name() //这里拿到的应该是main.main函数，如果需要拿到f()那么这里传递的应该是Caller(0)
	fmt.Println(funcName)
	fmt.Println(fine) // runtime_demo/main.go
	fmt.Println(line) // 12，是去找这个runtime.caller所执行的行数
}

func f1() {
	f()
}
func main() {
	f1() //注意这里的f1()套了一层，所以应该传2
	// pc, fine, line, ok := runtime.Caller(0) //注意这里传0是当前main函数里面执行的package,那比如说上面有个f1()函数，这里的Caller就需要传到1
	// if !ok {
	// 	fmt.Printf("runtime.Caller() failed\n")
	// 	return
	// }
	// fmt.Println(pc)
	// fmt.Println(fine) // runtime_demo/main.go
	// fmt.Println(line) // 12，是去找这个runtime.caller所执行的行数
}
