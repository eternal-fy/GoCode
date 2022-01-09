package main

import "fmt"

func main() {
	/*
		Go 语言中的数组是一种值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 `new()` 来创建： `var arr1 = new([5]int)`。
		那么这种方式和 `var arr2 [5]int` 的区别是什么呢？arr1 的类型是 `*[5]int`，而 arr2 的类型是 `[5]int`。
	*/
	//数组类型，当作形式参数传入，更改数据，值不会变
	arr1 := [4]int{5, 6, 7, 8}
	test1(arr1)
	fmt.Printf("%v", arr1) //5 6 7 8

	arr := []int{1, 2, 3, 4}
	test(arr)
	fmt.Printf("%v", arr) //2 3 4 5

}
func test1(arr [4]int) {
	for ix := 0; ix < len(arr); ix++ {
		arr[ix] += 1
	}
}
func test(arr []int) {
	for ix := 0; ix < len(arr); ix++ {
		arr[ix] += 1

	}
}
