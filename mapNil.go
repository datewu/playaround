package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["lol"] = 3
	m["ww"] = 6
	m["cf"] = 4
	fmt.Println(m)

	// m = nil
	m = make(map[string]int)
	m["new"] = 2
	m["lol"] = 0
	fmt.Println(m)
}
