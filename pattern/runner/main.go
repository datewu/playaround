package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"./runner"
)

const timeout = 16 * time.Second

func main() {
	fmt.Println("Starting work, guys")
	ctx, canncel := context.WithTimeout(context.Background(), timeout)
	defer canncel()

	r := runner.New(ctx)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Fatalln("Terminating due to timeout.")
		case runner.ErrInterrupt:
			log.Fatalln("Terminating due to interrupt.")
		default:
			log.Println(err)
		}
	}

	fmt.Println("Process End")
}

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Duration(2*id+2) * time.Second)
		fmt.Printf("Processor - Task #%d\n", id)
	}

}
