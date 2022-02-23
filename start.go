package main

func main() {
	//DatabaseConnect()
	//Handle()
	slice1 := make([]int, 2, 4)
	slice2 := append(slice1, 2, 3)
	for _, ans := range slice2 {
		print(ans, " ")
	}
	println(&slice1[0], &slice2[0]) //如果只添加2个，因为容量是4，加两个没有超过容量，所以不会重新分配内存，如果添加两个以上，就会重新分配，两个位置的地址值不一样

}
