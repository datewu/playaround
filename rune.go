package main

import (
	"fmt"
)

func main() {
	s := "我爱你"
	fmt.Printf("% xn, %[1]s\n", s)
	fmt.Println(len(s), s)

	b := []byte(s)
	fmt.Printf("% xn, %[1]s\n", b)
	fmt.Println(len(b), b)

	r := []rune(s)
	fmt.Printf("% xn, %[1]s\n", string(r))
	fmt.Println(len(r), r)
}
