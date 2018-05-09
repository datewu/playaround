package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	ctx, cannel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cannel()

	statusStream := make(chan opStatus)
	go doExpensiveOperation(ctx, statusStream)

	for {
		select {
		case status := <-statusStream:
			log.Println("greet >> read from status chan...")
			if status == opCanceled {
				log.Println("greet >> operation was cancelled")
				http.Error(w, "internal request was cancelled", http.StatusInternalServerError)
				return
			}
			if status == opTimeout {
				log.Println("greet >> operation timeout")
				http.Error(w, "internal request took too long ", http.StatusInternalServerError)
				return
			}
			log.Println("greet >> responding to client...")
			fmt.Fprintln(w, "hello gopher")
			return
		default:
			time.Sleep(time.Second)
			log.Println("Greetings are hard, Thinking...")
		}
	}

}

type opStatus int

const (
	opCompleted opStatus = iota
	opCanceled
	opTimeout
)

func doExpensiveOperation(ctx context.Context, ch chan<- opStatus) {
	log.Println("doExpensiveOperation >> called...")

	completionTimeout := time.After(4 * time.Second)

	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err == context.Canceled {
				ch <- opCanceled
				return
			} else if err == context.DeadlineExceeded {
				ch <- opTimeout
				return
			}
		case <-completionTimeout:
			log.Println("doExpensiveOperation >>  operation complete")
			ch <- opCompleted
			return

		default:
			log.Println("doExpensiveOperation >> slow operation...")
			time.Sleep(time.Second)
		}
	}

}
