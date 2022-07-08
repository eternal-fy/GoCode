package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.Buffer{}
	n, err := buffer.WriteString("hehe")
	if err != nil {
		panic(err)
	}
	println(n)
	fmt.Println(buffer.String())

	buf := make([]byte, 5)
	copy(buf[1:], "haha")
	fmt.Println(buf)

	arr := [...]int{1, 2, 3}
	fmt.Printf("%T", arr)
	fmt.Printf("%T", buf)

}
