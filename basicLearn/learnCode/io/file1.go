package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "d:/hehe.txt"
	outputFile := "d:/copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0) // 第三个参数权限  0   000  000 000 其中000分别为rwx   0开启1关闭  三个分别为owner group other
	if err != nil {
		panic(err.Error())
	}
}
