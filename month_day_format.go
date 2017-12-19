package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	for i := 0; i < 32; i++ {
		log.Println(i, now.Day(), now.Format("2006.01.02"))
		now = now.Add(time.Hour * 24)
	}
}
