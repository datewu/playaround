package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go func() {
			fmt.Println(i*10 + 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done")
}
