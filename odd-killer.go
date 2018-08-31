package main

import (
	"fmt"
)

func main() {
	//fmt.Println(luck(10000000000))
	fmt.Println(luck2(10000))
}

// binary tree
// the left-most && the-lowest element is the survivor
func luck2(size int) (survivor int) {
	survivor = 2
	for size > 2 {
		size /= 2
		survivor *= 2
	}
	return
}

func luck(size int) int {
	origin := make([]int, size)
	for i := 0; i < size; i++ {
		origin[i] = i + 1
	}
	remainer := even(origin)
	for len(remainer) > 1 {
		remainer = even(origin)
		origin = remainer
	}
	return remainer[0]
}

func even(a []int) (b []int) {
	b = make([]int, len(a)/2)

	for i := 0; i < len(b); i++ {
		b[i] = a[i*2+1]
	}
	return b
}
