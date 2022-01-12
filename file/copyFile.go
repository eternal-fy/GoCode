package main

import (
	"io"
	"os"
)

/**
复制文件夹
*/
func main() {
	//file, err := ioutil.ReadFile("E:/img/best.jpg")
	//
	//if err != nil {
	//
	//}
	//ioutil.WriteFile("e:/best.jpg", file, 0666)

	/*	open, err := os.Open("d:/file1")
		if err != nil {
			fmt.Printf("%v", open)
		}
		dir, err := open.ReadDir(-1)
		for _, d := range dir {
			fmt.Printf("%v\n", d)
		}*/
	/*rootPath := "d:/file1"
	dir, err := ioutil.ReadDir(rootPath)
	if err != nil {

	}

	for _, d := range dir {
		fmt.Printf("%v", d)
	}
	*/
	//复制文件夹代码
	CopyDir("d:/file1", "d:/file2")
}

func CopyDir(src, des string) {
	//srcFile, _ := os.Open(src)
	_, err := os.Stat(des)
	//如果des的文件存在，则全部删除
	if err == nil {
		os.RemoveAll(des)
	}
	os.Mkdir(des, 0777)
	Bl(src, des)
}

//遍历
func Bl(src, des string) {
	srcFile, _ := os.Open(src)
	readdir, _ := srcFile.Readdir(0)
	for _, d := range readdir {
		newSrc := src + "/" + d.Name()
		newDes := des + "/" + d.Name()
		if d.IsDir() {
			os.Mkdir(newDes, 0777)
			Bl(newSrc, newDes)
		} else {
			srcNewFile, _ := os.Open(newSrc)
			desNewFile, _ := os.Create(newDes)
			io.Copy(srcNewFile, desNewFile)
		}
	}
}
