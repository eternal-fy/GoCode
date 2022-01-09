package main

import (
	"flag"
	"runtime"
)

func main() {
	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
}
