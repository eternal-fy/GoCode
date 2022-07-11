package main

type Animal interface { //某个结构体类型  将接口里的方法都实现了就可以使用多态
	speak()
	walk()
}
type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}
func (p *Person) Name() string {
	return p.name
}
func (p *Person) speak() {
	println("person  speak  ", p.name)
}
func (p *Person) walk() {
	println("person  walk   ", p.name)
}
func main() {
	var p Animal = &Person{name: "ld"}
	p.walk()
	p.speak()
}
