package main

import "fmt"

type in interface {
	lol(string) string
}

type lala struct{}

func (l *lala) lol(s string) string {
	return "hello dota"
}

func main() {
	var l in = lala{}
	fmt.Println(l.lol("dd"))
}
