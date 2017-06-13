package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var str string
	if len(os.Args) > 2 {
		str = os.Args[1]

	}
	m := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}
	fmt.Println(m)
	var result int
	for i, r := range str {
		n := len(str) - i - 1
		if ra, ok := m[r]; ok {
			result += int(ra * int(math.Pow10(n)))
		} else {
			panic("no such character")
		}
	}
	fmt.Println(result)
}
