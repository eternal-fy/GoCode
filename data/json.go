package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string //命名首字符要大写，否则无法Encode
}

func main() {
	p := Person{"ld"}
	ans, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("error")
	}
	fmt.Println(string(ans))

}
