package main

func main() {
	str := "abcd"
	chars := []byte(str)
	for i := 0; i < len(chars); i++ {
		println(i, chars[i])
		chars[i] = 'a'
	}
	println(string(chars))

}
