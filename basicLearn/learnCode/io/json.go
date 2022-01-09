package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func (v *VCard) String() string {
	return "{FirstName:" + v.FirstName + " LastName:" + v.LastName + " Remark:" + v.Remark + "}"
}
func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	outFile, _ := os.Open("d:/json.txt")
	outFile.Write(js)
	var v VCard
	json.Unmarshal(js, &v)
	println("---------")
	fmt.Printf("%v", vc)
	println()
	fmt.Printf("%v", v)

	// using an encoder:

}
