package main

import (
	"fmt"
	"os"
)

func main() {
	path := "d:/demo.txt"
	write(path)
	read(path)
}

func read(path string) {
	file, err := os.Open(path)
	if err != nil {
		HandleError(err)
	}
	buf := make([]byte, 2048)
	file.Read(buf)
	fmt.Println(string(buf))

}
func write(path string) {

	file, err := create(path)
	defer file.Close()
	if err != nil {
		HandleError(err)
	}

	file.WriteString("haha")
}
func HandleError(err error) {
	fmt.Println("err--", err)
	os.Exit(1)
}
func create(name string) (file *os.File, err error) {
	file, err = os.Create(name)
	return
}
