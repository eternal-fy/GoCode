package main

import (
	"fmt"
)

/*
主要是接口类型指向继承接口的对象时，判断被指向的对象是什么类型
*/
type shape interface {
	getArea() float32
}
type square struct {
	height, width float32
}
type circle struct {
	radius float32
}

type triangle struct {
	a, b, c float32
}

func (s *square) getArea() float32 {
	return s.width * s.height
}

func (s *circle) getArea() float32 {
	return 3.14 * s.radius * s.radius
}

func (s *circle) String() string {
	strRadius := fmt.Sprintf("%f", s.radius)
	return "circle:{radius:" + strRadius + "}"
}

func (s *square) String() string {
	strHeight := fmt.Sprintf("%f", s.height)
	strWidth := fmt.Sprintf("%f", s.width)
	return "square:{height:" + strHeight + " width:" + strWidth + "}"
}

func main() {
	var point shape
	s := new(square)
	s.height = 10.1
	s.width = 10.2
	point = s
	judgeType(point)

	r := new(circle)
	r.radius = 12
	point = r
	judgeType(point)
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

	//验证一个struct是否实现了某个接口

	cc1 := &circle{12}
	aa1 := &triangle{1, 2, 3}
	ss1 := &square{1, 2}
	var cc, aa, ss interface{}
	cc = cc1
	aa = aa1
	ss = ss1

	if sv, ok := cc.(shape); ok {
		fmt.Println("circle implements shape--", sv)
	}
	if sv, ok := aa.(shape); ok {
		fmt.Println("triangle implements shape--", sv)
	} else {
		fmt.Printf("triangle no implements")
	}
	if sv, ok := ss.(shape); ok {
		fmt.Println("square implements shape--", sv)
	}

}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

func judgeType(point shape) {
	switch t := point.(type) {
	case *square:
		fmt.Println("--square--")
	case *circle:
		fmt.Println("--circle--")
	case nil:
		fmt.Println("---nil---")
	default:
		fmt.Println("---others---", t, "--")
	}
}
