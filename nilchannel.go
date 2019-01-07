package main

func main() {
	var c chan int
	<-c
	close(c)
}
