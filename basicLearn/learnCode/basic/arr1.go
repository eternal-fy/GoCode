package main

func main() {
	arr := make([]int, 5)
	for i := range arr {
		arr[i] = i + 1
	}
	change(arr)
	for i := range arr {
		println(arr[i])
	}
}
func change(arr []int) {
	for i := range arr {
		arr[i] += 1
	}
}
