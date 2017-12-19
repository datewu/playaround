package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		p = make(chan struct{})
		n int64
	)

	go func() {
		p <- struct{}{}
	}()

	go func() {
		for {
			<-p
			n++
		}
	}()
	go func() {
		for {
			p <- struct{}{}
			//	n++
		}
	}()

	t := time.Now()
	time.Sleep(time.Second)
	fmt.Println(n, time.Since(t))
}
