package main

import (
	"Gone/rpc/demo"
	"time"
)

func main() {
	go demo.Server()
	time.Sleep(1 * time.Second)
	demo.Client()
}
