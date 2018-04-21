package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	if err := checkRecoverInGoroutine(); err != nil { // must go before normal check
		fmt.Println("this is a panic attact recover in goroutine", err)
		return
	}

	if err := checkRecover(); err != nil {
		fmt.Println("this is a panic attact recover in function call", err)
		return
	}

	time.Sleep(15 * time.Second)
	fmt.Println("no panic in last 18 second ")
}

func checkRecoverInGoroutine() (err error) {
	log.Println("check Goroutine")
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("send on close channel")
		}
	}()
	go func() {
		fmt.Println("go routine will panic in 3 second")
		time.Sleep(3 * time.Second)
		panic("i am panic")
	}()
	return nil
}

func checkRecover() (err error) {
	log.Println("check normal")
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("panic normal")
		}
	}()
	time.Sleep(3 * time.Second)
	panic("i am panic")
	return nil
}
