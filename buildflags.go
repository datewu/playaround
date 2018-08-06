//go run -ldflags="-X main.dota=sdfsdf -X main.lol=989" main.go
package main

import "fmt"

const dota = "good"

var lol = "v5"

func main() {
	fmt.Println("const ", dota)
	fmt.Println("variable", lol)
}
