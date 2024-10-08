package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var s string
	fmt.Print("input ur message")
	fmt.Scanln(&s)
	fmt.Printf("你输入的是%s\n", s)

}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	//直到读到了\n换行符为止
	fmt.Printf("您输入的是%s", s)
}

func main() {

}
