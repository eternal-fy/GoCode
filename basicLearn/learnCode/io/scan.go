package main

import "fmt"

func main() {

	var (
		firstName, lastName string
	)
	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName, lastName)

	//open, _ := os.Open("d:/demo.txt")
	//inputReader := bufio.NewReader(open)
	//fmt.Println("Please enter some input: ")
	//for {
	//	input, flag, err := inputReader.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(input), "---", flag, "---", err)
	//
	//}

}
