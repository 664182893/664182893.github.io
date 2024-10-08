package main

import "fmt"

type cat struct{}

type dog struct{}

type person struct {
}

type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}

func (c cat) speak() {
	println("喵喵喵！")
}

func (d dog) speak() {
	println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("aaa")
}
func da(x speaker) {
	x.speak() //挨打了就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var e1 person
	da(c1)
	da(d1)
	da(e1)

}
