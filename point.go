package main

import "fmt"

type lol struct {
	count   int
	sss, ss int
}

func main() {
	han := lol{count: 7}
	fmt.Println(han)
	c := han.count
	c = 9
	fmt.Println(han, c)

	p := &(han.count)
	*p = 10
	fmt.Println(han)

	m := newlol()
	m.count = 9
	fmt.Println(m)

	// 	newlol().count = 4  compile error
	newlolp().count = 4
}

func newlol() lol {
	return lol{}
}

func newlolp() *lol {
	return new(lol)
}
