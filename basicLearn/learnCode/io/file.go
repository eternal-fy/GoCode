package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, e := os.Open("d:\\hehe.txt")
	defer file.Close()
	if e != nil {
		fmt.Println("open fail!\n")
		return
	}
	input := bufio.NewReader(file)
	for {
		in, err := input.ReadString('\n')
		fmt.Print(in)
		if io.EOF == err {
			return

		}

	}
}
