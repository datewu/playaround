package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	if os.Args[1] != "lol" {
		log.Println("dota")
	}

}
