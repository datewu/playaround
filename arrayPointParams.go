package main

import "fmt"

type a struct {
	name, hobby string
}

func main() {
	m := []*a{
		&a{name: "jobn", hobby: "gamej"},
		&a{name: "smil", hobby: "cf"},
		&a{name: "mayr", hobby: "gcf"},
		&a{name: "tao", hobby: "doat"},
		&a{name: "haot", hobby: "atdo"},
	}
	pretty(m)
	fmt.Println("vim-go", m)
	modify(m)
	fmt.Println("after modifyo", m)
	pretty(m)
}

func modify(data []*a) {
	for i, v := range data {
		if i == 1 {
			continue
		}
		v.name = "lol"
	}
}

func pretty(data []*a) {
	for i, v := range data {
		fmt.Println(i, v.name, v.hobby)
	}
}
