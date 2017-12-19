package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Comming coundown")
	tick := time.Tick(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		select {
		case <-tick:
		case <-abort:
			log.Println("Launch aborted")
			return
		}
	}
	launch()
}
func launch() {
	fmt.Println("WoW")
}
