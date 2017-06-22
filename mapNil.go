package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["lol"] = 3
	m["dota"] = 4
	m["cf"] = 6
	fmt.Println(m)
	//m = nil
	m = make(map[string]int)
	m["t"] = 12
	fmt.Println(m)
}
