package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

/*
文件处理的几种方式
*/

func main() {
	src := "E:/my.txt"
	des := "E:/mysec.txt"
	Method3(src, des)
}

func Method1(src, des string) {
	file, err := os.Open(src) //open只读模式，Create可以读写和修改O_RDWR|O_CREATE|O_TRUNC   都是调用的OpenFile函数，这个设置更加精细
	HandleErro(err)
	buf := make([]byte, 18)
	f, e := os.Create(des)
	HandleErro(e)
	ok := true
	for ok {
		size, e := file.Read(buf)
		HandleErro(e)
		ok = e != io.EOF
		println(string(buf[:size]))
		f.Write(buf[:size])
	}
}

func Method2(src, des string) {
	file, _ := os.Open(src)
	reader := bufio.NewReader(file)
	desfile, _ := os.Create(des)
	writer := bufio.NewWriter(desfile)

	for ok := true; ok; {
		readString, e := reader.ReadString('\n')
		println(readString)
		writer.WriteString(readString)
		writer.Flush() //一定要刷新，否则缓存没入本地内存
		ok = e != io.EOF
	}

}

func Method3(src, des string) {
	filebytes, _ := ioutil.ReadFile(src)
	println(string(filebytes))
	ioutil.WriteFile(des, filebytes, 0666)
}

func HandleErro(err error) {
	if err != nil {
		println(err)
	}

}
