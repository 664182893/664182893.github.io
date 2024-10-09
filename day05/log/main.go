package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写文件
// writefile
func writefilebufio() {
	file, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("input error! err:%v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 1; i++ {
		writer.WriteString("writer")
	}
	writer.Flush()
}

func writefileioutil() {
	str := "wangsss"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("open file failed,err :%v", err)
		return
	}
}

func writefile() {
	file, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err :%v", err)
		return
	}
	// fmt.Println(file)
	//write
	file.Write([]byte("吱吱吱!"))
	file.WriteString("萨瓦迪卡！")
	file.Close()
}

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	//直到读到了\n换行符为止
	fmt.Printf("您输入的是%s", s)
	file, _ := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	fmt.Fprintln(file, s)
}
