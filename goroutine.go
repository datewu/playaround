package main

import "fmt"

func main() {
	var c = make(chan int)
	go func() {
		c <- 0
	}()

	for {
		go func() {
			m := <-c
			m++
			c <- m
			fmt.Println(m)
		}()
	}
}
