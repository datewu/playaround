package main

import (
	"fmt"
	"log"
)

func main() {
	sli := make([]int, 2)
	sli = append(sli, 998)
	for _, v := range sli {
		fmt.Println(v)
	}
	sli = nil
	sli = append(sli, 14)
	log.Println("after nil")

	for _, v := range sli {
		fmt.Println(v)
	}
	sli = nil
}
