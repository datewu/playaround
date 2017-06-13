package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go echo(i)
		go func() {
			echo(i)
		}()

	}
	time.Sleep(3 * time.Second)
}

func echo(i int) {
	//time.Sleep(1 * time.Second)
	fmt.Println(i)
}
