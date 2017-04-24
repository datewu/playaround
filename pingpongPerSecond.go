package main

import (
	"fmt"
	"time"
)

type l struct{}

var (
	pp    = make(chan string)
	count int64
)

func main() {
	// ping
	go func() {
		for {
			pp <- "ping"
		}
	}()

	go func() {
		for {
			<-pp
			count++
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(count)
}
