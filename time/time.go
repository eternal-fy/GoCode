package main

import (
	"fmt"
	"time"
)

func main() {
	after := time.After(1 * time.Second)

	timer := time.NewTimer(2 * time.Second)
	println("---+++++")
	m := <-timer.C
	fmt.Printf("%v\n", m)
	t := <-after
	fmt.Println(t)

	println("-----")
	timer1 := time.NewTimer(1 * time.Second)
	timer1.Reset(5 * time.Second)
	ans := <-timer1.C
	fmt.Printf("%v\n", ans)

	ticker := time.NewTicker(2 * time.Second)
	for {
		<-ticker.C
		fmt.Println("滴滴！！")
	}

}
