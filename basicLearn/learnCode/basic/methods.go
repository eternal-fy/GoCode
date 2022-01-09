package main

import (
	"fmt"
	"strconv"
)

type node struct {
	num1 int
	num2 int
}
type element int

func main() {
	var e element
	e = 10
	e.show()

	no := &node{1, 2}
	println(no.add())

	no1 := &node{1, 2}
	fmt.Printf("%v", no1)

}
func (n *node) String() string { //创建String()string方法,标准化输出的时候会自动调用String()方法
	return "node:{num1:" + strconv.Itoa(n.num1) + " num2:" + strconv.Itoa(n.num2) + "}"
}

func (e element) show() {
	println(e, "---")

}
func (n *node) add() int {
	return n.num1 + n.num2
}
