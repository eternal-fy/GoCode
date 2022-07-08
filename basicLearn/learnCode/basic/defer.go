package main

import (
	"fmt"
)

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
	/*
		e b
		in b
		e a
		in a
		l a
		l b
	*/
	i := 0
	for ; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
	i++
	// 11 11 11 .....
}
