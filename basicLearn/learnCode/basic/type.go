package main

import "fmt"

func main() {
	var decimal float32
	decimal = 10.222222222
	println(decimal)
	fmt.Printf("%.3f\n", decimal)
	ans := int(decimal)
	println(ans)
	println("complex--------")
	c := 10 + 20i
	fmt.Printf("%v\n", c)
	var re, im float32
	re = 10
	im = 20
	cc := complex(re, im) //re,im必须为float类型
	fmt.Printf("%v", cc)

}
