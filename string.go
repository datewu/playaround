package main

import "fmt"

func main() {
	s := []byte("lolo")
	s[0] = 's'
	fmt.Println(string(s))
}
