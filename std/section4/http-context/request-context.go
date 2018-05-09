package main

import (
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

	completeAfter := time.After(5 * time.Second)
	ctx := r.Context()

	for {
		select {
		case <-completeAfter:
			fmt.Fprintf(w, "hello gopher")
			return
		case <-ctx.Done():
			log.Println("Context error:", ctx.Err())
			http.Error(w, "you cannel request", http.StatusRequestTimeout)
			return
		default:
			time.Sleep(time.Second)
			log.Println("Greetings are hard, Thinking...")
		}
	}

}
