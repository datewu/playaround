package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	var s string
	if len(os.Args) < 2 {
		log.Println("Not enough arguments")
		return
	}
	s = os.Args[1]

	m := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	}

	fmt.Println(m)
	var result int
	for i, r := range s {
		n := len(s) - i - 1
		if ra, ok := m[r]; ok {
			result += int(ra * int(math.Pow10(n)))
		} else {
			panic("no such character")
		}
	}
	fmt.Println(s, "Done:", result)

}
