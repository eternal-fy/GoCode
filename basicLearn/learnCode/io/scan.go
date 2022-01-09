package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	/*
		var(
		firstName,lastName string
		)
		fmt.Scanln(&firstName,&lastName)
		fmt.Println(firstName,lastName)
	*/

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("input :  %s  ---", input)
	}

}
