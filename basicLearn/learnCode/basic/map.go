package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["ld"] = 12
	map1["ld1"] = 22
	fmt.Printf("%v %v %v", map1["ld"], map1["he"], len(map1))
	if val, ok := map1["ld"]; ok {
		fmt.Printf("--- %v  %v", val, ok)
	}
	delete(map1, "ld")
	println(len(map1))

	println("-------------------")

	map2 := make(map[int]float32)
	for i := 0; i < 5; i++ {
		map2[i] = float32(i)
	}

	for key, val := range map2 {
		println(key, "-", val)

	}

}
