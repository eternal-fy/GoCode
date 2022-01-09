package main

func main() {

	ans := WrapAdd(1, 2, Add)
	println(ans)

}
func Add(a, b int) (c int) {

	c = a + b
	return
}

func WrapAdd(a, b int, f func(int, int) int) int {
	return f(a, b)
}
