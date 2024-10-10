package main

import (
	"fmt"
	"io"
	"os"
)

func f2() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	// fileObj.Seek(3, 0) //光标移动到首字符往后的第三个字符 windows的\n为\r\n
	// var ret string
	// ret = "8"
	// str := []byte(ret)
	// n, err := fileObj.WriteAt(str, 3)
	if err != nil {
		fmt.Printf("error :err%v", err)
		return
	}
	defer fileObj.Close()
	//因为没有办法直接在文件中间插入内容，所以要借助一个临时文件。
	tmpFile, err1 := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error :err%v", err1)
		return
	}
	defer tmpFile.Close()
	//读取文件写入临时文件
	var ret1 [1]byte
	n, err2 := fileObj.Read(ret1[:])
	if err != nil {
		fmt.Printf("error :err%v", err2)
		return
	}
	//写入临时文件
	tmpFile.Write(ret1[:n])
	//再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)
	//再传入剩余的源文件内容
	var x [1024]byte
	//i, err := fileObj.Read(x[:1024]); err !=（应该改为==） nil;(注意这里当没有1024byte数的时候err==nil)
	for {
		i, err := fileObj.Read(x[:1024])
		if err == io.EOF {
			fmt.Println(i)
			tmpFile.Write(x[:i])
			break
		}
		if err != nil {
			fmt.Printf("error :err%v", err2)
			return
		}
		tmpFile.Write(x[:i])
	}
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")
	// fmt.Println(string(ret[:n]))
}

func main() {
	f2()
}
