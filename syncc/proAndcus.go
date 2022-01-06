package syncc

import (
	"fmt"
	"time"
)

var producer, customer int

func init() {
	producer, customer = 0, 0
}

func PCstart() {
	go produce()
	go custome()
}
func produce() {
	println("---")
	for {
		producer++
		fmt.Println("producer-----", producer)
		time.Sleep(100 * time.Millisecond)
	}

}
func custome() {
	println("+++")
	for {
		if customer < producer {
			customer++
			fmt.Println("SELL----", customer)
			time.Sleep(100 * time.Millisecond)
		}
	}

}
