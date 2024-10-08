package main

import "fmt"

type santala struct {
	name string
}
type baoma struct {
	name string
}
type car interface {
	run()
}

func drive(x car) {
	x.run()
}
func (a santala) run() {
	fmt.Printf("%s70迈")
}

func (b baoma) run() {
	fmt.Printf("%s70迈")
}
func main() {
	var a = santala{name: "hi"}
	drive(a)
}
