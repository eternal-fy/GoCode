package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 12.34
	ans := reflect.TypeOf(a)
	fmt.Println(ans, reflect.Float64)
	ans1 := reflect.ValueOf(a)                           //ans1类型变成了reflect.Value
	fmt.Println(ans1, reflect.TypeOf(ans1), ans1.Kind()) //reflect.Value类型的变量，可以通过kind()方法返回变量的类型，例如：float64

	fmt.Println("------------")
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	fmt.Println("------------")
	num := 12.34
	val1 := reflect.ValueOf(&num) //一定要加&
	fmt.Println(val1.CanSet())
	point1 := val1.Elem()
	fmt.Println(point1.CanSet())
	point1.SetFloat(34.11)
	fmt.Println(num)
}
