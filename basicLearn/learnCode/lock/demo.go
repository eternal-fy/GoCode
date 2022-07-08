package main

import (
	"fmt"
	"sync"
	"time"
)

type Info struct {
	lock sync.Mutex
}

var info Info
var pfood int
var cfood int

func main() {
	pfood = 0
	cfood = 0
	info = Info{}
	go Produce()
	go Custome()
	select {}

}

func Produce() {
	for true {
		info.lock.Lock()
		pfood++
		time.Sleep(10)
		fmt.Println("food produce ------------- ", pfood)
		info.lock.Unlock()
	}

}

func Custome() {
	for true {
		if cfood < pfood {
			info.lock.Lock()
			cfood++
			time.Sleep(10)
			fmt.Println("food sell ---- ", cfood)
			info.lock.Unlock()
		}
	}

}
