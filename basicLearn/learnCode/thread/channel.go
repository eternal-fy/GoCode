package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	fmt.Println("---------")
	go suck(stream)
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(1e8)
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
