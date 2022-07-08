package main

import "fmt"

func main() {
	//如果切片容量够，就自己添加，否则扩容产生新的切片
	arr := []int{1, 2, 3, 4}
	ints := append(arr, 4, 5)
	arr[0] = 0
	fmt.Printf("%x  %x", arr, ints)
	//[0 2 3 4]  [1 2 3 4 4 5]

	arr1 := make([]int, 4, 6)
	for i := 0; i < 4; i++ {
		arr1[i] = i + 1
	}

	ints1 := append(arr1, 4, 5)
	arr1[0] = 0
	fmt.Printf("%x  %x", arr1, ints1)
	//[0 2 3 4]  [0 2 3 4 4 5]
}
