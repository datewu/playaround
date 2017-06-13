package main

import "fmt"

func main() {
	s := "中古海峡"
	fmt.Printf("% xn, %[1]s\n", s)
	fmt.Println(len(s))

	b := []byte(s)
	fmt.Printf("% xn, %[1]s\n", b)
	fmt.Println(len(b))

	r := []rune(s)
	fmt.Printf("% xn, %[1]s\n", string(r))
	fmt.Println(len(r))
}
