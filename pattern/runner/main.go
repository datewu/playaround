package main

import (
	"log"
	"time"

	"./runner"
)

func main() {
	log.Println("starting work.")

	r := runner.New(timeout)

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

const timeout = 30 * time.Second

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
