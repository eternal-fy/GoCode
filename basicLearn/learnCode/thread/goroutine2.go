package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e12)
}

func sendData(ch chan string) {
	ch <- "Washington"
	time.Sleep(1e9)
	ch <- "Tripoli"
	time.Sleep(1e9)
	ch <- "London"
	time.Sleep(1e9)
	ch <- "Beijing"
	time.Sleep(1e9)
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string

	for {
		time.Sleep(5e9)
		input = <-ch //管道接收者和发送者都准备就绪才可以传递，否则阻塞
		fmt.Printf("%s ", input)
	}
}
