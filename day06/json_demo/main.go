package main

import (
	"encoding/json"
	"fmt"
)

//json

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"周林","age":9000}`
	var p person
	json.Unmarshal([]byte(str), &p) //注意这里传值要对应传递结构体类型，根据结构体的key去对应匹配value
	fmt.Println(p.Name, p.Age)
}
