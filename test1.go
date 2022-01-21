package main

import "fmt"

func main() {

	var b [10000]int
	var n, i int
	fmt.Scanf("%d", &n)
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &b[i])
	}
	for i = 0; i < n; i++ {
		fmt.Println("%d", b[i])
	}

}
