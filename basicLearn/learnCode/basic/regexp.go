package main

import (
	"fmt"
	"regexp"
)

func main() {
	str2 := "alkdfja0932uorj4irhi499 90.23aerl;aeri3"
	regexw := "[0-9]+\\.[0-9]"
	if ok, _ := regexp.Match(regexw, []byte(str2)); ok {
		fmt.Println("matches!")
	}
	pat, _ := regexp.Compile(regexw)
	str1 := pat.ReplaceAllString(str2, "----")
	fmt.Println(str1)

	str := "abcd123ba"
	regex := "[0-9]+"
	if ok, _ := regexp.Match(regex, []byte(str)); ok {
		fmt.Println("match")
	}
	compile, _ := regexp.Compile(regex)
	str = compile.ReplaceAllString(str, "-----")
	fmt.Printf(str)
}
