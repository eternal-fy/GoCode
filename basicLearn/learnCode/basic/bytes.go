package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	fmt.Printf("%T", buffer)
	buffer.WriteString("abc")
	fmt.Printf("%v\n", buffer.String())
	println()

	buffer1 := new(bytes.Buffer)
	fmt.Printf("%T", buffer1)
	buffer1.WriteString("def")
	fmt.Printf("%v", buffer1.String())

}
