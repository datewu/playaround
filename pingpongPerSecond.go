package main

import (
	"fmt"
	"time"
)

func main() {
	var pp = make(chan int)
	go func() {
		pp <- 0
	}()

	t := time.Now()
	var count int
	go func() {
		for {
			m := <-pp
			m++
			pp <- m
			count = m
		}
	}()

	go func() {
		for {
			n := <-pp
			n++
			pp <- n
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(count, time.Since(t))
}
