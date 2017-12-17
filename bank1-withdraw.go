package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	deposits = make(chan int) // send amount to deposit
	balances = make(chan int) // receive balance
	withdraw = make(chan wd)
)

type wd struct {
	amount int
	c      chan bool
}

// Withdraw amount
func Withdraw(amount int) {
	newWd := wd{amount, make(chan bool)}
	withdraw <- newWd
	if <-newWd.c {
		fmt.Println("success withdrwa:", amount, "remain:", Balance())
		return
	}
	log.Println("failed during withdraw:", amount, "remain:", Balance())
}

// Deposit amount
func Deposit(amount int) {
	deposits <- amount
}

// Balance get the balance
func Balance() int {
	return <-balances
}

func teller() {
	var b int // b is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			b += amount
		case balances <- b:
		case w := <-withdraw:
			if w.amount <= b {
				b -= w.amount
				w.c <- true
			} else {
				w.c <- false
			}
		}

	}
}

func main() {
	go teller()

	log.Println("Get blance", Balance())
	Deposit(150)
	log.Println("Deposit 150", Balance())

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(amount int) {
			Withdraw(3 + 10*amount)
			time.Sleep(200 * time.Millisecond)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
