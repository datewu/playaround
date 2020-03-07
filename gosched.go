package main

import (
	"fmt"
	//	"runtime"
	"time"
)

func main() {
	fmt.Println("vim-go")
	go func() {
		fmt.Println("am i run going to run?")
		time.Sleep(3 * time.Second)
		fmt.Println("am i run too?")
	}()
	//	runtime.Gosched()
	fmt.Println("bye")
}
