package main

import (
	"fmt"
)

type aa struct {
	m, f func(n int) error
}

func main() {
	a := &aa{}
	a.m = func(i int) error {
		fmt.Println("run in a.m", i)
		return nil
	}
	//	fmt.Println("run a.f()")
	//a.f(12)

	fmt.Println("run a.f()")
	a.m(12)
}
