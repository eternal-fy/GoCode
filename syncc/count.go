package syncc

import (
	"fmt"
	"sync"
	"time"
)

var count int
var lock sync.Mutex

func init() {
	count = 0

}
func CountStart() {
	go add1()
	go add2()
	time.Sleep(2 * time.Second)
	fmt.Println("count------------", count)
}
func add1() {
	for i := 0; i < 3000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}

}
func add2() {
	for i := 0; i < 3000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
}
