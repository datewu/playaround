package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	for i := 1; i < 32; i++ {
		log.Println(i, now.Day(), now.Format("2006.01.02"))
		now = now.Add(24 * time.Hour)
	}
}
