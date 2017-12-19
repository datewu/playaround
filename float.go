package main

import (
	"fmt"
)

func main() {
	var f float32 = 1 << 24
	fmt.Println(f, f == f+1)

	var m float32 = 16777215
	fmt.Println(m, m+2 == m+1)

	fmt.Println(f - m)

}
