package main

import (
	"fmt"
	"log"
	"time"
)

var chStream = make(chan struct{})

func main() {
	lol()
	<-chStream
	fmt.Println("done")
}

func lol() {
	fmt.Println("enting lol")
	time.Sleep(time.Second)
	fmt.Println("enting goroutine")
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("exit goroutine")
		chStream <- struct{}{}
	}()
	time.Sleep(time.Second)
	fmt.Println("exit lol")
}
