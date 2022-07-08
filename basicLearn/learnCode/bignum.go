package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := big.NewRat(7, 3)
	n := big.NewRat(3, 7)
	mul := m.Mul(m, n)
	f, _ := mul.Float64()
	fmt.Printf("%v", f)
}
