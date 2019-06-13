package main

import "fmt"

type Person interface {
	greet()
}

type person struct {
	name string
	age  int
}

func (p *person) greet() {
	fmt.Println("H1, my name is", p.name)
}

func newPerson(name string, age int) Person {
	return &person{
		name,
		age,
	}
}

func main() {
	p := newPerson("lol", 998)
	p.greet()
	//	fmt.Println(p.name, p.age)
}
