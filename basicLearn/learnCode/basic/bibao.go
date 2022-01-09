package main

import "fmt"

func f() (ret int) {
	println(&ret, "-")
	defer func() {
		println(&ret, "--")
		ret++
	}()
	return 1
}

func fun1() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}
func main() {
	func1 := func(a, b int) int { return a + b }
	ans1 := func1(1, 2)
	println(ans1)
	addr := f()
	fmt.Println(&addr, "---")
	println("---------------")
	funAns := fun1()
	ans := funAns(1, 2)
	println(ans)
}
