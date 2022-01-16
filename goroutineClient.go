package main

import (
	"fmt"
	"os"
)

func main() {
	open, err := os.OpenFile("e:/now.txt", os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error")
	}
	open.Write([]byte("hehe\n"))
}
