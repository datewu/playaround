package main

import (
	"log"
	"time"
)

func main() {
	t := time.AfterFunc(100*time.Second,
		func() {
			log.Println("exit func")
		})
	log.Println("going to for loop")
	for i := 0; i < 5; i++ {
		log.Println("in for loop", i)
		time.Sleep(5 * time.Second)
		log.Println("resetting timmer 10s", t.Reset(10*time.Second))
		//log.Println("resetting timmer 10s", t.Reset(10*time.Second), t.Stop())
	}
	time.Sleep(5 * time.Minute)

	//	runtime.Goexit()
	log.Println("exiting main")
}
