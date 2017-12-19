package main

import (
	"fmt"
)

func main() {
	var m []int
	m = make([]int, 100)
	m[6] = 4
	fmt.Println(m[:40])
}
