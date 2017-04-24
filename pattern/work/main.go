package main

import (
	"log"
	"sync"
	"time"

	"./work"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"rack",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, v := range names {
			np := namePrinter{
				v,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	log.Println("dispatch DONE", "LOL")
	wg.Wait()
	p.Shutdown()
}
