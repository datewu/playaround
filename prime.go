package main

import (
	"fmt"
)

func main() {
	source := make(chan int)
	go generate(source)
	for i := 1; i < 10; i++ {
		prime := <-source
		fmt.Println(i, prime)
		filterSource := make(chan int)
		go filter(source, filterSource, prime)
		source = filterSource
	}
}

func generate(output chan<- int) {
	for i := 2; ; i++ {
		output <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		number := <-in
		if number%prime != 0 {
			out <- number
		}
	}
}
