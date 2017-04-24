// Package bank provides a concurrency-safe bank with one account.
package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Get balance", Balance())
	Deposit(150)
	log.Println("Deposit 150", Balance())

	go Withdraw(10)
	go Withdraw(10)
	go Withdraw(10)
	go Withdraw(10)
	go Withdraw(10)
	go Withdraw(10)
	go Withdraw(20)
	go Withdraw(30)
	go Withdraw(40)
	go Withdraw(50)
	go Withdraw(60)
	go Withdraw(80)
	time.Sleep(3 * time.Second)

}

var (
	deposits = make(chan int) // send amount to deposit
	balances = make(chan int) // receive balance
	withdraw = make(chan wa)
)

type wa struct {
	amount int
	c      chan bool
}

func Withdraw(amount int) {
	newWa := wa{amount, make(chan bool)}
	withdraw <- newWa
	if <-newWa.c {
		log.Println("success withdraw:", amount, "reamin:", Balance())
		return
	}
	log.Println("failed durning withdraw:", amount, "reamin:", Balance())
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// nothing
		case w := <-withdraw:
			if w.amount <= balance {
				balance -= w.amount
				w.c <- true
				break // NOTE: cannot use return
			}
			w.c <- false
		}
	}
}

func init() {
	go teller()
}
