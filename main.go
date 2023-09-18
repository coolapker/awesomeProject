package main

import (
	"fmt"
	"github.com/coolapker/NBlog"
)

func main() {

	c := make(map[int]int)
	fmt.Printf("%v\n", c)
	fmt.Println(len(c))

}
func changeName(cat *Cat) {
	fmt.Printf("inner  %p\n", &cat)
}

func change(s []int) {
	(s)[0] = 100
}

type Animal interface {
	run()
}
type Cat struct {
	Name string
}

func (c *Cat) run() {
	fmt.Println(c.Name, " running___")
	c.Name = "cat"
}
