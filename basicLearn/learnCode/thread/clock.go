package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e9)
	boom := time.After(5e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			time.Sleep(5e7)
		}
	}
}
