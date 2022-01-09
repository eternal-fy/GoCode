package main

import (
	"fmt"

	"os"
)

func main() {
	inputFile, _ := os.Open("d:/game.ico")
	buf := make([]byte, 1024*100)
	n, _ := inputFile.Read(buf)
	fmt.Println(n)
	outputFile, _ := os.Open("d:/gamecopy.ico")
	z, _ := outputFile.Write(buf)
	fmt.Println(z)

}
