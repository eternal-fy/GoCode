package main

import (
	"fmt"
)

/*
主要是接口类型指向继承接口的对象时，判断被指向的对象是什么类型
*/

type triangle struct {
	a, b, c float32
}

func (s *Circle) String() string {
	strRadius := fmt.Sprintf("%f", s.radius)
	return "Circle:{radius:" + strRadius + "}"
}

func (s *Square) String() string {
	strHeight := fmt.Sprintf("%f", s.height)
	strWidth := fmt.Sprintf("%f", s.width)
	return "Square:{height:" + strHeight + " width:" + strWidth + "}"
}

func main() {
	var point Shape
	s := new(Square)
	s.height = 10.1
	s.width = 10.2
	point = s
	judgeType(point)

	r := new(Circle)
	r.radius = 12
	point = r
	judgeType(point)
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

	//验证一个struct是否实现了某个接口

	cc1 := &Circle{12}
	aa1 := &triangle{1, 2, 3}
	ss1 := &Square{1, 2}
	var cc, aa, ss interface{}
	cc = cc1
	aa = aa1
	ss = ss1

	if sv, ok := cc.(Shape); ok {
		fmt.Println("Circle implements Shape--", sv)
	}
	if sv, ok := aa.(Shape); ok {
		fmt.Println("triangle implements Shape--", sv)
	} else {
		fmt.Printf("triangle no implements")
	}
	if sv, ok := ss.(Shape); ok {
		fmt.Println("Square implements Shape--", sv)
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

func judgeType(point Shape) {
	switch t := point.(type) {
	case *Square:
		fmt.Println("--Square--")
	case *Circle:
		fmt.Println("--Circle--")
	case nil:
		fmt.Println("---nil---")
	default:
		fmt.Println("---others---", t, "--")
	}
}
