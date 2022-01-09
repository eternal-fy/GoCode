package main

import (
	"os"
)

func main() {
	f, _ := os.Open("d:/hehe.txt")
	buf := make([]byte, 1024)
	read, err := f.Read(buf)
	if err == nil {
		print("---")
		os.Stdout.WriteString(string(buf[0:read]))
	}

}
