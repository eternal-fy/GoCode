package main

import "fmt"

func main() {
	num := 10
	ChangeD(num)
}
func ChangeD(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T-int-", v)
	case int32:
		fmt.Printf("%T-int32-", v)

	}
	fmt.Printf("%v", num)
}
