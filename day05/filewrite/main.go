package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件

// func filewrite() {

// }

func main() {
	// file, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Printf("open file failed,err :%v", err)
	// 	return
	// }
	// // fmt.Println(file)
	// //write
	// file.Write([]byte("吱吱吱!"))
	// file.WriteString("萨瓦迪卡！")
	// file.Close()
	readfile()
	iouitlfile()
}

func readfile() {
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file failed,err :%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 1; i++ {
		writer.WriteString("hello")
	}
	writer.Flush() //将缓存的数据写入文件
}

func iouitlfile() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	//这部分会清空原文件内容，然后再写入数据
	if err != nil {
		fmt.Printf("open file failed,err :%v", err)
		return
	}
}

//上述通过writestring去实现
//另外或者可以通过bufio.newWriter去实现
//再者，可以通过ioutil.Writefile
