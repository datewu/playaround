package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// set up channel on which to send signal notifications.
	// we must use a buffered channel or risk missig the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	signal.Notify(c)
	fmt.Println("Got signal: ", <-c)
	fmt.Println("Done")
}
