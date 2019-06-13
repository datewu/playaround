package main

import (
	"log"
	"time"
)

func main() {
	t := 3 * time.Second
	timer := time.NewTimer(t)
	for {
		select {
		case ts := <-timer.C:
			log.Println(ts)
			timer.Reset(t)
		default:
		}

	}
}
