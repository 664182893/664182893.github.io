package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 通过ioutil去读取
func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file error,err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

// 通过ioutil去读取
func readFromFileByIoutil() {
	file, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file error: err%v\n", err)
		return
	}
	fmt.Println(file)
}

// 使用bufio去读取文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("Open file failed,err:%v", err)
	}
	defer fileObj.Close()
	//上述记得关闭文件
	reader := bufio.NewReader(fileObj)
	//创建一个用来从文件中读内容的对象
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	// fileObj, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Printf("open file failed,err:%v", err)
	// 	return
	// }
	// defer fileObj.Close()
	// var tmp [128]byte
	// //读文件
	// for {
	// 	n, err := fileObj.Read(tmp[:])
	// 	if err == io.EOF {
	// 		fmt.Println("读完了")
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("read failed!:%v", err)
	// 		return
	// 	}
	// 	fmt.Printf("读了%d个字节", n)
	// 	fmt.Println(string(tmp[:n]))
	// 	if n < 128 {
	// 		return
	// 	}
	// }
	// readFromFilebyBufio()
	readFromFileByIoutil()
}
