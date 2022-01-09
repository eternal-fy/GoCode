package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(123456789)
	b := big.NewInt(2) //1/2
	fmt.Printf("%v  %v  %v", b.Mul(b, b), a.Mul(a, a), a.Mul(a, b))

}
