package main

import "fmt"

func main() {
	c := make(chan bool)
	close(c)
	for i := 0; i < 100; i++ {
		select {
		case <-c:
			fmt.Println(i)
		default:
			close(c)
		}
	}
	fmt.Println("vim-go")
}
