package main

import (
	"io"
	"os"
)

func main() {
	src := "d:/game.ico"
	srcFile, _ := os.Open(src)
	des := "d:/gamecopy.ico"
	desFile, _ := os.Create(des)
	io.Copy(desFile, srcFile)

}
