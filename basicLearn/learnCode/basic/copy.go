package main

import "fmt"

func main() {
	arr := make([]int, 2, 10)
	arr1 := []int{1, 2, 3}
	n := copy(arr, arr1) //(to,from)

	arr1[0] = 10 //说明是上面的是深拷贝[1 2] [10 2 3]  2
	fmt.Printf("%v %v  %v", arr, arr1, n)

}
