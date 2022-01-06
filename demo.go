package main

import (
	"fmt"
)

func main() {
	arr := []interface{}{1, 3.23, true, "hehe"}

	justify(1, 3.23, true, "hehe")
	fmt.Println("---")
	justify(arr...)
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)

}

func justify(etype ...interface{}) {
	for _, m := range etype {
		switch v := m.(type) {
		case int:
			fmt.Println("int=====", v)
		case string:
			fmt.Println("stringz=====", v)
		case float64:
			fmt.Println("float64=====", v)
		case bool:
			fmt.Println("bool=====", v)
		default:
			fmt.Println("other type----")
		}

	}

}
