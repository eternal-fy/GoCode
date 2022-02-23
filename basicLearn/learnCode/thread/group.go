package main

import "sync"

var group = sync.WaitGroup{}

func main() {

	for i := 0; i < 10; i++ {
		group.Add(1)
		go speak(i)
	}
	group.Wait()

}
func speak(msg int) {
	println(msg)
	group.Done()
}
