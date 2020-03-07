package main

import "fmt"
import "time"

func main() {
	go func() {
		fmt.Println(block())
	}()
	fmt.Println("vim-go")
}

func block() int {
	fmt.Println("in block")
	time.Sleep(5 * time.Second)
	fmt.Println("out block")
	return 222
}
