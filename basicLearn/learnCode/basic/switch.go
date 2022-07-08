package main

import "fmt"

func main() {
	value := 'a'
	switch {
	case value == 'b':
		println(value, "match")
	case value == 'a':
		println(value, "match")
	}

	switch value {
	case ' ':
		println("null")
	case 'b', 'c', 'a':
		println("match")

	}

	ans := 18
	ans1 := interface{}(ans)
	i := 10
	switch v := ans1.(type) {
	case bool:
		fmt.Printf("Param #%v is a bool\n", v)
	case float64:
		fmt.Printf("Param #%d is a float64\n", i)
	case int, int64:
		fmt.Printf("Param #%d is a int\n", i)
	case nil:
		fmt.Printf("Param #%d is a nil\n", i)
	case string:
		fmt.Printf("Param #%d is a string\n", i)
	default:
		fmt.Printf("Param #%d is unknown\n", i)
	}

}
