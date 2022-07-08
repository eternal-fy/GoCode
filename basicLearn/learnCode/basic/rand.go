package main

import (
	"math/rand"
	"time"
)

func main() {
	s := int64(time.Now().Nanosecond())
	rand.Seed(s)
	arr := make([]int, 10)
	for i := 0; i < 100000; i++ {
		n := rand.Intn(10)
		arr[n]++
		//z := rand.Float32()
	}
	for _, item := range arr {
		print(item, " ")
	}
}
