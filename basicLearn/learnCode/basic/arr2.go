package main

func main() {
	ints := []int{1: 10, 2, 3}
	Change(ints)
	for i := range ints {
		println(ints[i])
	}
}

func Change(arr []int) {
	for i := range arr {
		arr[i] += 1
	}
}
