package main

import (
	"fmt"
	"sync"
	"time"

	"./work"
)

var names = []string{
	"steve blown",
	"bob green",
	"marry ch",
	"the reverse flash",
	"rack dave",
	"willion well",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	time.Sleep(time.Second)
	fmt.Println(n.name)
}

func main() {
	p := work.New(200)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, v := range names {
			np := namePrinter{v}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	fmt.Println("dispatch Done, LOL")
	wg.Wait()
	p.Shutdown()
	fmt.Println("tasks Done")
}
