package main

import "fmt"

func main() {
	var m []int
	m = make([]int, 1000)
	m[9] = 1
	fmt.Println(m[:99])
}
