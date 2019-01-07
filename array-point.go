package main

import (
	"fmt"
)

func main() {
	m := [4]int{2: 8}
	fmt.Println("befre zero", m)

	zero(&m)
	fmt.Println("befre zero", m)
}

func zero(p *[4]int) {
	// for i := range p {
	// 	(*p)[i] = 0
	// 	p[i] = 0
	// }
	*p = [4]int{}
}
