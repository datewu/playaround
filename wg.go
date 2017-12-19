package main

import "sync"
import "fmt"
import "time"

func main() {
	var (
		a  = []int{23, 5, 6, 70, 173, 200, 180, 170, 190, 45, 62, 98}
		wg sync.WaitGroup
	)

	wg.Add(len(a))
	for _, v := range a {
		go worker(&wg, v)
	}
	wg.Wait()
	fmt.Println("done")
}

func worker(wg *sync.WaitGroup, during int) {
	defer wg.Done()
	time.Sleep(time.Duration(10*during) * time.Millisecond)
	fmt.Println(during)

}
