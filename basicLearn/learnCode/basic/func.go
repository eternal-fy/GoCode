package main

import "fmt"

func main() {

	ans := show(2, "hehe", "aha")
	showArr(ans...)
}
func show(len int, str ...string) []string {
	fmt.Printf("%T", len)
	fmt.Printf("%T", str)
	for index, val := range str {
		println(index, val)
	}
	return str
}
func showArr(str ...string) {
	for i, val := range str {
		println(i, val)
	}
}
