package main

import (
	"fmt"
	"time"
)

func main() {
	//var c = make(chan int)
	var c chan int // a send or receive on a nil chan block forever
	go func() {
		c <- 3
	}()

	//	fmt.Println(<-c)
	time.Sleep(3 * time.Second)
	fmt.Println("lol")
}
