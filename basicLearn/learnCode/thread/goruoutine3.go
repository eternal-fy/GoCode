package main

import "fmt"

func main() {
	buf := 100
	ch := make(chan string, buf)
	ch <- "hehe"
	ch <- "haha"
	var m string
	m = <-ch
	fmt.Printf("%v", m) //hehe
	m = <-ch
	fmt.Printf("%v", m) //haha

}
