package main

import (
	"fmt"
	"time"
)

func main() {
	var p = make(chan int)
	var q = make(chan int)
	var n int
	go func() {
		p <- 1
	}()

	t := time.Now()

	go func() {
		for {

			<-p
			n++
			q <- 1
		}
	}()
	go func() {
		for {
			<-q
			n++
			p <- 1
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(n, time.Since(t))
	/*
		var tick = time.Tick(time.Second)
		time.Sleep(990 * time.Millisecond)
		for {
			select {
			case <-p:
				go func() {
					q <- 1
					n++
				}()

			case <-q:
				go func() {
					p <- 1
					n++
				}()
			case <-tick:
				fmt.Println(n, time.Since(t))
				os.Exit(0)
			}
		}
	*/
}
