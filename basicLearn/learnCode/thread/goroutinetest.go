package main

func main() {
	ch := make(chan string)
	ch <- "hehe"

	println("-----")

}
