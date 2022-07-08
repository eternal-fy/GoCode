package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	println("program----")

}

func main() {

	var goos string = runtime.GOOS
	println("GOOS:")
	println(goos)
	path := os.Getenv("PATH")
	println("PATH:")
	println(path)
	a, b := 3, 4 //并行赋值
	println(a, b)
	a, b = b, a
	println(a, b)
	a++ //只能这样单独使用
	println(a, b)
	c := 1234.456789101
	fmt.Printf("%.3e", c) //1.234e+03
	/*变量声明方法-------------------------------
	  var (
	  	a = 15
	  	b = false
	  	str = "Go says hello to the world!"
	  	numShips = 50
	  	city string
	  )



	  var(
	  	point *int
	  	num1 int
	  	flag bool
	  	decimal float32
	  	//num2 int
	  	//str string
	  	)
	  	println("init values")
	  	println(point,num1,flag,decimal)
	  	num1=10
	  	point=&num1
	  	num1=12
	  	flag=true
	  	if true{
	  		println(1)
	  	}else{
	  		println(0)
	  	}
	  	println(num1,flag,*point)
	*/

}
