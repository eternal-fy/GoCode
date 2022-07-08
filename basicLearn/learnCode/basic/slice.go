package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(len(arr), cap(arr))
	slice := arr[:]
	fmt.Println(len(slice), cap(slice))

	slice1 := make([]int, 5, 10)
	slice1 = slice1[:cap(slice1)-3] //扩展切片，最高扩展为cap（）
	fmt.Printf("%v\n", slice1)

	val1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	sval := val1[3:4] //切片的容量为从左边取得第一个到数组最后一个数，前面的不计数
	fmt.Println(len(sval), cap(sval), sval)

}
