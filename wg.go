package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var a = []int{2, 34, 5, 6, 7, 8}

func main() {
	go log.Println(http.ListenAndServe(":80", nil))
	go wgTest()
	log.Println(http.ListenAndServe(":8099", nil))
}

func wgTest() {
	time.Sleep(time.Second)
	wg.Add(len(a))
	for _, v := range a {
		go worker(v)
	}
	wg.Wait()
	log.Println("exit wgTest")
}

func worker(p int) {
	defer wg.Done()
	time.Sleep(1800 * time.Millisecond)
	fmt.Println(p)
}
