package main

import "fmt"

type em struct {
	ID   int
	name string
}

func main() {
	employID(34).name = "dota"
	fmt.Println("vim-go", employID(3))
}

func employID(id int) *em {
	return &em{98, "lol"}
}
