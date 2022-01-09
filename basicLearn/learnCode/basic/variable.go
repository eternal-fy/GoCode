package main

import (
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
