package main

import "fmt"

const PI float32 = 3.14

type Shape interface {
	getArea() float32
}

type Square struct {
	height, width float32
}
type Circle struct {
	radius float32
}

func (s Square) getArea() float32 {
	return s.height * s.width
}

func (c Circle) getArea() float32 {
	return c.radius * c.radius * PI
}
func main() { //接口定义后就是指针，只能只想继承者的方法，无法更改变量属性
	var point Shape
	point1 := new(Square)
	point = point1
	point1.height = 9
	point1.width = 8
	fmt.Println(point.getArea())

}
