package main

import "io/ioutil"

func main() {
	file, err := ioutil.ReadFile("E:/img/best.jpg")

	if err != nil {
	}
	ioutil.WriteFile("e:/best.jpg", file, 0666)
}
