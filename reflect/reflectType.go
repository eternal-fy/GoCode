package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

func main() {
	a := 12.3
	of := reflect.TypeOf(a)
	p := Person{"ld"}
	fmt.Println(of)
	fmt.Println(reflect.TypeOf(p))
}
