package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("not enough arguments ")
		return
	}
	fmt.Println(os.Args[1])
	if os.Args[1] != "lol" {
		log.Println("dota")
	}
}
