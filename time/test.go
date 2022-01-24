package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	date := now.AddDate(0, 1, 1)
	fmt.Printf("%v  ", now)
	yearDay := date.YearDay()
	day := now.Day()
	println(yearDay, day)
	ticker := time.NewTicker(10 * time.Second)
	didi := <-ticker.C
	fmt.Printf("%v", didi)

}
