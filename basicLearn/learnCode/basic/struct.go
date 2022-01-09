package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	num1 float32 "tag1--num" //后面的是属性的标签，可以通过showTag方法进行获取
	flag bool    "tag2---flag"
}

type Node1 struct {
	num1 int
	Node
	int
}

func main() {
	var num Node
	fmt.Println(num)
	num1 := new(Node)
	fmt.Println(num1)
	num2 := Node{12.2, true}
	fmt.Println(num2)
	num3 := &Node{10.0, true}
	fmt.Println(num3)
	f := func(a, b int) int { return a + b }
	ans := f(10, 11)
	fmt.Println(ans)

	showTag(num2)

	//ans2:=pac.getInstance(2,"ld",23)   getInstance，小写标识是私有的方法，无法调用
	//fmt.Println(ans2)

	fmt.Println("--------")

	conode := new(Node1)
	conode.num1 = 10 //如果两个结构体有同名的数据，则需要通过匿名的方式进行访问conode.Node.num1
	conode.flag = true
	fmt.Println(conode)

}
func showTag(node Node) {
	//field.Name 是变量名称   field.Tag是tag信息
	for i := 0; i < 2; i++ {
		ix := i
		ty := reflect.TypeOf(node)
		field := ty.Field(ix)
		fmt.Println(field.Tag)
	}

}
