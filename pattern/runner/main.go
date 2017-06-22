package main

import (
	"context"
	"log"
	"time"

	"./runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("starting work.")
	ctx, canncel := context.WithTimeout(context.Background(), timeout)
	defer canncel()

	r := runner.New(ctx)

	r.Add(createTask(), createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Fatalln("Terminating due to timeout.")
		case runner.ErrInterrupt:
			log.Fatalln("Terminating duro interrupt.")
		}
	}
	log.Println("Process ended。")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
