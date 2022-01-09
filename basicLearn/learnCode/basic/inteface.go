package main

import "fmt"

const PI float32 = 3.14

type shape interface {
	getArea() float32
}

type square struct {
	height, width float32
}
type circle struct {
	radius float32
}

func (s square) getArea() float32 {
	return s.height * s.width
}

func (c circle) getArea() float32 {
	return c.radius * c.radius * PI
}
func main() { //接口定义后就是指针，只能只想继承者的方法，无法更改变量属性
	var point shape
	point1 := new(square)
	point = point1
	point1.height = 9
	point1.width = 8
	fmt.Println(point.getArea())

}
